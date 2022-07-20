package main

import (
	"context"
	"fmt"
	"github.com/camunda/zeebe/clients/go/v8/pkg/zbc"
	"os"
)

const ZeebeAddr = "172.30.52.239:26500"

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

	// deploy process
	ctx := context.Background()
	response, err := zbClient.NewDeployResourceCommand().AddResourceFile("order-process.bpmn").Send(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.String())

	createInstanceCount := 10
	for i := 0; i < createInstanceCount; i++ {

		// create a new process instance
		variables := make(map[string]interface{})
		variables["orderId"] = i + 1000

		request, err := zbClient.NewCreateInstanceCommand().BPMNProcessId("order-process").LatestVersion().VariablesFromMap(variables)
		if err != nil {
			panic(err)
		}

		result, err := request.Send(ctx)
		if err != nil {
			panic(err)
		}
		fmt.Println(result.String())

	}

}
