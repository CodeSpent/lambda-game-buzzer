package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

func handleRequest(ctx context.Context, event events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	rc := event.RequestContext
	switch rk := rc.RouteKey; rk {
	case "$connect":
		break
	case "$disconnect":
		break
	case "$roundStart":
		break
	case "$buzz":
		break
	case "$default":
		break
	default:
		log.Fatalf("Unknown RouteKey %v", rk)
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handleRequest)
}
