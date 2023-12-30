package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	// "github.com/aws/aws-sdk-go/aws"
    // "github.com/aws/aws-sdk-go/aws/session"
    // "github.com/aws/aws-sdk-go/service/apigatewaymanagementapi"
)

type MyEvent struct {
	Name string `json:"name"`
	Message string `json:"message"`
	Action string `json:"action"`
	RequestContext struct {
        ConnectionID string `json:"connectionId"`
    } `json:"requestContext"`
}

func HandleRequest(ctx context.Context, event *MyEvent) (*string, error) {
	if event == nil {
		return nil, fmt.Errorf("received nil event")
	}

	connectionID := event.RequestContext.ConnectionID
	fmt.Sprint(connectionID)

	message := fmt.Sprintf("%s : %s. Action: %s", event.Name, event.Message, event.Action)
	return &message, nil
}

func main() {
	lambda.Start(HandleRequest)
}