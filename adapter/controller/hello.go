package controller

import (
	"clean-serverless-book-sample/registry"
	"clean-serverless-book-sample/usecase"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

type PostHelloRequest struct {
	Name string `json:"name"`
}

type HelloMessageResponse struct {
	Message string `json:"message"`
}

var ValidateHelloMessageSettings = []*ValidatorSetting{
	{ArgName: "name", ValidateTags: "required"},
}

func PostHello(request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	validErr := ValidateBody(request.Body, ValidateHelloMessageSettings)
	if validErr != nil {
		return Response400(validErr)
	}

	var req PostHelloRequest
	err := json.Unmarshal([]byte(request.Body), &req)
	if err != nil {
		return Response500(err)
	}

	h := registry.GetFactory().BuildCreateHelloMessage()
	res, err := h.Execute(&usecase.CreateHelloMessageRequest{
		Name: req.Name,
	})
	if err != nil {
		return Response500(err)
	}

	return Response200(&HelloMessageResponse{
		Message: res.Message,
	})
}
