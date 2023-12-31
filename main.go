package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/apigatewaymanagementapi"
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
    }


	sess := session.Must(session.NewSession())
	if(sess == nil){
		fmt.Print("SESSION IS NULL")
	}
	endpoint := "https://" + event.RequestContext.DomainName + "/" + event.RequestContext.Stage
	apiGwManagementApi := apigatewaymanagementapi.New(sess, aws.NewConfig().WithEndpoint(endpoint))
	message2 := "Hello, react client!"

	fmt.Printf("ENDPOINT: %s\n", endpoint)
	fmt.Printf("CONNECTIONID %s\n", event.RequestContext.ConnectionID)
	_, err := apiGwManagementApi.PostToConnection(&apigatewaymanagementapi.PostToConnectionInput{
		ConnectionId: aws.String(event.RequestContext.ConnectionID),
		Data:         []byte(message2),
	})
	if err != nil {
		fmt.Printf("Error sending message: %s\n", err.Error())
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

func handle_connect(ctx context.Context, event *MyEvent) (Response, error){
	// Handle connect
	connectionID := event.RequestContext.ConnectionID
	message := fmt.Sprintf("%s connected", connectionID)
	response := Response{
		StatusCode: 200, // HTTP status code
		Body:       message,
	}
	return response, nil
}

func handle_disconnect(ctx context.Context, event *MyEvent) (Response, error){
	// Handle connect
	connectionID := event.RequestContext.ConnectionID
	message := fmt.Sprintf("%s connected", connectionID)
	response := Response{
		StatusCode: 200, // HTTP status code
		Body:       message,
	}
	return response, nil
}