package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
)

func handle_disconnect(ctx context.Context, event *MyEvent) (Response, error){
	// Handle connect
	connectionID := event.RequestContext.ConnectionID
	message := fmt.Sprintf("%s connected", connectionID)

	err := removeConnectionIDFromDynamoDB(ctx, event.RequestContext.ConnectionID)
    if (err != nil){
        fmt.Printf("ERROR ADDING ITEM TO DYNAMODB: %s\n", err)
    }

	response := Response{
		StatusCode: 200, // HTTP status code
		Body:       message,
	}
	return response, nil
}

func removeConnectionIDFromDynamoDB(ctx context.Context, connectionID string) error {
    // Create a new AWS session
    sess := session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
    }))

    // Create a DynamoDB client
    svc := dynamodb.New(sess)

    // Define the DeleteItem input
    input := &dynamodb.DeleteItemInput{
        TableName: aws.String("websocket-connectionIds"),
        Key: map[string]*dynamodb.AttributeValue{
            "ConnectionID": {
                S: aws.String(connectionID),
            },
        },
    }

    // Delete the item from the DynamoDB table
    _, err := svc.DeleteItemWithContext(ctx, input)
    if err != nil {
        return fmt.Errorf("failed to delete item from DynamoDB: %w", err)
    }

    fmt.Printf("Successfully removed connection ID %s from DynamoDB\n", connectionID)
    return nil
}
