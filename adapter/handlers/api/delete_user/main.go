package main

import (
	"clean-serverless-book-sample/adapter/controller"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return controller.DeleteUser(request), nil
}

func main() {
	lambda.Start(handler)
}
