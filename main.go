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
		RouteKey  string `json:"routeKey"`
    } `json:"requestContext"`
}

type Response struct {
    StatusCode int    `json:"statusCode"`
    Body       string `json:"body"`
}

func HandleRequest(ctx context.Context, event *MyEvent) (Response, error) {
    switch event.RequestContext.RouteKey {
    case "$connect":
        // Handle connect
		connectionID := event.RequestContext.ConnectionID
		message := fmt.Sprint("%s connected",connectionID)
		response := Response{
			StatusCode: 200, // HTTP status code
			Body:       message,
		}
		return response, nil
    case "$disconnect":
        // Handle disconnect
		connectionID := event.RequestContext.ConnectionID
		message := fmt.Sprint("%s disconnected",connectionID)
		response := Response{
			StatusCode: 200, // HTTP status code
			Body:       message,
		}
		return response, nil
    default:
        // Handle default message
    }

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