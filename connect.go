package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type WebSocketConnection struct {
    ConnectionID string `json:"ConnectionID" dynamodbav:"ConnectionID"`
}

func handle_connect(ctx context.Context, event *MyEvent) (Response, error){
	// Handle connect
	connectionID := event.RequestContext.ConnectionID
	message := fmt.Sprintf("%s connected \n", connectionID)

	err := addConnectionIDToDynamoDB(ctx, event.RequestContext.ConnectionID)
    if (err != nil){
        fmt.Printf("ERROR ADDING ITEM TO DYNAMODB: %s\n", err)
    }
	response := Response{
		StatusCode: 200, // HTTP status code
		Body:       message,
	}
	return response, err
}


func addConnectionIDToDynamoDB(ctx context.Context, connectionID string) error {
    // Create a new AWS session
    sess := session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
    }))
    fmt.Print("Created new session")
    // Create a DynamoDB client
    svc := dynamodb.New(sess)
    fmt.Print("Created new client")
    // Create an instance of the WebSocketConnection struct
    item := WebSocketConnection{
        ConnectionID: connectionID,
    }
    fmt.Print("mapped item")
    // Marshal the Go struct into a DynamoDB attribute value map
    av, err := dynamodbattribute.MarshalMap(item)
    if err != nil {
        return fmt.Errorf("got error marshalling new item: %w", err)
    }

    // Define the PutItem input
    input := &dynamodb.PutItemInput{
        Item:      av,
        TableName: aws.String("websocket-connectionIds"),
    }

    // Put the item into the DynamoDB table
    _, err = svc.PutItemWithContext(ctx, input)
    if err != nil {
        return fmt.Errorf("failed to put item in DynamoDB: %w", err)
    }

    fmt.Printf("Successfully added connection ID %s to DynamoDB\n", connectionID)
    return nil
}