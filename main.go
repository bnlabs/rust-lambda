package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
	Message string `json:"message"`
	Action string `json:"action"`
	RequestContext struct {
        ConnectionID	string `json:"connectionId"`
		RouteKey 		string `json:"routeKey"`
		DomainName   	string `json:"domainName"`
        Stage        	string `json:"stage"`
    } `json:"requestContext"`
}

type Response struct {
    StatusCode int    `json:"statusCode"`
    Body       string `json:"body"`
}

func HandleRequest(ctx context.Context, event *MyEvent) (Response, error) {
    switch event.RequestContext.RouteKey {
    case "$connect":
		response, error := handle_connect(ctx, event)
		return response, error
    case "$disconnect":
		response, error := handle_disconnect(ctx, event)
		return response, error
    default:
        // Handle default message
		response, error := handle_default(ctx, event)
		return response, error
    }
}

func main() {
	lambda.Start(HandleRequest)
}