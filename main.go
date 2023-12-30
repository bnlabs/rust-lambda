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

type Response struct {
    StatusCode int    `json:"statusCode"`
    Body       string `json:"body"`
}

func HandleRequest(ctx context.Context, event *MyEvent) (Response, error) {
	connectionID := event.RequestContext.ConnectionID
	fmt.Sprint(connectionID)

	message := fmt.Sprintf("%s: %s. Action: %s", event.Name, event.Message, event.Action)

	response := Response{
        StatusCode: 200, // HTTP status code
        Body:       message,
    }
	return response, nil
}

func main() {
	lambda.Start(HandleRequest)
}