package main

import (
	"context"
	"fmt"
	"github.com/camunda/zeebe/clients/go/v8/pkg/entities"
	"github.com/camunda/zeebe/clients/go/v8/pkg/worker"
	"github.com/camunda/zeebe/clients/go/v8/pkg/zbc"
	"log"
	"os"
	"time"
)

const ZeebeAddr = "10.12.7.29:26500"

var readyClose = make(chan struct{})

func main() {
	gatewayAddr := os.Getenv("ZEEBE_ADDRESS")
	plainText := false

	if gatewayAddr == "" {
		gatewayAddr = ZeebeAddr
		plainText = true
	}

	zbClient, err := zbc.NewClient(&zbc.ClientConfig{
		GatewayAddress:         gatewayAddr,
		UsePlaintextConnection: plainText,
	})

	if err != nil {
		panic(err)
	}

	fmt.Printf("_________________ Connected _______________________")

	// deploy process
	ctx := context.Background()
	response, err := zbClient.NewDeployResourceCommand().AddResourceFile("order-process.bpmn").Send(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.String())

	// create a new process instance
	variables := make(map[string]interface{})
	variables["orderId"] = "31243"

	request, err := zbClient.NewCreateInstanceCommand().BPMNProcessId("order-process").LatestVersion().VariablesFromMap(variables)
	if err != nil {
		panic(err)
	}

	result, err := request.Send(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.String())

	jobWorker := zbClient.NewJobWorker().JobType("payment-service").Handler(handleJob).
		Name("Worker1").
		MaxJobsActive(10).
		PollInterval(time.Second * 3).
		PollThreshold(2).
		Open()

	<-readyClose
	jobWorker.Close()
	jobWorker.AwaitClose()
}

func handleJob(client worker.JobClient, job entities.Job) {
	jobKey := job.GetKey()

	headers, err := job.GetCustomHeadersAsMap()
	if err != nil {
		// failed to handle job as we require the custom job headers
		failJob(client, job)
		return
	}

	variables, err := job.GetVariablesAsMap()
	if err != nil {
		// failed to handle job as we require the variables
		failJob(client, job)
		return
	}

	variables["totalPrice"] = 46.50
	request, err := client.NewCompleteJobCommand().JobKey(jobKey).VariablesFromMap(variables)
	if err != nil {
		// failed to set the updated variables
		failJob(client, job)
		return
	}

	log.Println("Complete job", jobKey, "of type", job.Type)
	log.Println("Processing order:", variables["orderId"])
	log.Println("Collect money using payment method:", headers["method"])

	ctx := context.Background()
	_, err = request.Send(ctx)
	if err != nil {
		panic(err)
	}

	log.Println("Successfully completed job")
	close(readyClose)
}

func failJob(client worker.JobClient, job entities.Job) {
	fmt.Printf("_________________ Error _______________________")

	log.Println("Failed to complete job", job.GetKey())

	ctx := context.Background()
	_, err := client.NewFailJobCommand().JobKey(job.GetKey()).Retries(job.Retries - 1).Send(ctx)
	if err != nil {
		panic(err)
	}
}
