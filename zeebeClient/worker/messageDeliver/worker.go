package main

import (
	"context"
	"fmt"
	"github.com/camunda/zeebe/clients/go/v8/pkg/zbc"
	"os"
)

const ZeebeAddr = "172.30.52.239:26500"

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

	variables := make(map[string]interface{})
	variables["message"] = "This is test message"
	message, err := zbClient.NewPublishMessageCommand().MessageName("getMessageFromOut").CorrelationKey("123456").VariablesFromMap(variables)

	ctx := context.Background()
	response, err := message.Send(ctx)

	fmt.Printf(response.String())
}
