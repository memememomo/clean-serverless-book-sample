package usecase

type ICreateHelloMesessage interface {
	Execute(req *CreateHelloMessageRequest) (*CreateHelloMessageResponse, error)
}

type CreateHelloMessageRequest struct {
	Name string
}

type CreateHelloMessageResponse struct {
	Message string
}
