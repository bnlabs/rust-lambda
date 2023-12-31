package main

import (
	"context"
	"fmt"
)

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