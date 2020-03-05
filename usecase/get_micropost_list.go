package usecase

import "clean-serverless-book-sample/domain"

type IGetMicropostList interface {
	Execute(req *GetMicropostListRequest) (*GetMicropostListResponse, error)
}

type GetMicropostListRequest struct {
	UserID uint64
}

type GetMicropostListResponse struct {
	Microposts []*domain.MicropostModel
}
