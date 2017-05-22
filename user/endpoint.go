package user

import (
	"github.com/go-kit/kit/endpoint"
	"context"
)

type userRequest struct {
	username string
	password string
}

type userResponse struct {
	msg string
	err error `json:"error,omitempty"`
}

func makeRegisterEndpoints(s Service) endpoint.Endpoint {
	return func(ctx context.Context,request interface{})(interface{},error) {
		req := request.(userRequest)
		_,err :=s.Register(req.username,req.password)
		if err!=nil {
			return &userResponse{msg:"failure",err:err}
		}
		return &userResponse{msg:"sucess",err:nil}
	}
}