package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/apigatewaymanagementapi"
)

func handle_default(ctx context.Context, event *MyEvent) (Response, error){
	sess := session.Must(session.NewSession())
	if(sess == nil){
		fmt.Print("SESSION IS NULL")
	}
	endpoint := "https://" + event.RequestContext.DomainName + "/" + event.RequestContext.Stage
	apiGwManagementApi := apigatewaymanagementapi.New(sess, aws.NewConfig().WithEndpoint(endpoint))
	message := "Hello, react client!"

	fmt.Printf("ENDPOINT: %s\n", endpoint)
	fmt.Printf("CONNECTIONID %s\n", event.RequestContext.ConnectionID)
	_, err := apiGwManagementApi.PostToConnection(&apigatewaymanagementapi.PostToConnectionInput{
		ConnectionId: aws.String(event.RequestContext.ConnectionID),
		Data:         []byte(message),
	})

	response := Response{
		StatusCode: 200,
		Body: message,
	}
	return response, err
}