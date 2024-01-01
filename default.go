package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/apigatewaymanagementapi"
	"encoding/json"
)

type data struct{
    Name    string `json:"name"`
    Message string `json:"message"`
}

func handle_default(ctx context.Context, event *MyEvent) (Response, error){
	fmt.Printf("Received event: %+v\n", event)
	sess := session.Must(session.NewSession())
	if(sess == nil){
		fmt.Print("SESSION IS NULL")
	}
	endpoint := "https://" + event.RequestContext.DomainName + "/" + event.RequestContext.Stage
	apiGwManagementApi := apigatewaymanagementapi.New(sess, aws.NewConfig().WithEndpoint(endpoint))
	responseData := data{
		Name: event.Name,
		Message: event.Message,
	}

	responseJSON, err := json.Marshal(responseData)
	if err != nil {
		fmt.Printf("error serializing response data %s", err.Error())
		responseError := Response{
			StatusCode: 400,
			Body: "error serializing responseData",
		}
		return responseError,err
	}

	_, err = apiGwManagementApi.PostToConnection(&apigatewaymanagementapi.PostToConnectionInput{
		ConnectionId: aws.String(event.RequestContext.ConnectionID),
		Data:         responseJSON,
	})

	response := Response{
		StatusCode: 200,
		Body: "",
	}
	return response, err
}