package interactor

import (
	"clean-serverless-book-sample/usecase"
	"fmt"
)

type CreateHelloMessage struct {
}

func NewCreateHelloMessage() *CreateHelloMessage {
	return &CreateHelloMessage{}
}

func (c *CreateHelloMessage) Execute(req *usecase.CreateHelloMessageRequest) (*usecase.CreateHelloMessageResponse, error) {
	msg := fmt.Sprintf("Hello, %s!", req.Name)
	return &usecase.CreateHelloMessageResponse{Message: msg}, nil
}
