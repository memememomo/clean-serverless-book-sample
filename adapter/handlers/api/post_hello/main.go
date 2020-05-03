package main

import (
	"clean-serverless-book-sample/adapter/controller"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handlers(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return controller.PostHello(request), nil
}

func main() {
	lambda.Start(handlers)
}
