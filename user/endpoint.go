package user

import (
	"github.com/go-kit/kit/endpoint"
	"context"
)

type userRequest struct {
	username string
	password string
	newPass string
}

type userResponse struct {
	Msg string
	Err string `json:"error,omitempty"`
}

func makeRegisterEndpoints(s Service) endpoint.Endpoint {
	return func(ctx context.Context,request interface{})(interface{},error) {
		req := request.(*userRequest)
		_,err :=s.Register(req.username,req.password)
		if err!=nil {
			return &userResponse{Msg:"failure", Err:err.Error()},nil
		}
		return &userResponse{Msg:"sucess", Err:""},nil
	}
}

func makeLoginEndpoints(s Service) endpoint.Endpoint {
	return func(ctx context.Context,request interface{})(interface{},error) {
		req := request.(*userRequest)
		exists,err :=s.Login(req.username,req.password)
		if err!=nil {
			return &userResponse{Msg:"failure", Err:err.Error()},nil
		}
		if !exists{
			return &userResponse{Msg:"failure", Err:"user does not exist"},nil
		}
		return &userResponse{Msg:"sucess", Err:""},nil
	}
}

func makeChangePasswordEndpoints(s Service) endpoint.Endpoint {
	return func(ctx context.Context,request interface{})(interface{},error) {
		req := request.(*userRequest)
		success,err :=s.ChangePassword(req.username,req.password,req.newPass)
		if err!=nil {
			return &userResponse{Msg:"failure", Err:err.Error()},nil
		}
		if !success {
			return &userResponse{Msg:"failure", Err:err.Error()},nil
		}
		return &userResponse{Msg:"sucess", Err:""},nil
	}
}

func makeDeleteEndpoints(s Service) endpoint.Endpoint {
	return func(ctx context.Context,request interface{})(interface{},error) {
		req := request.(*userRequest)
		_,err :=s.Delete(req.username,req.password)
		if err!=nil {
			return &userResponse{Msg:"failure", Err:err.Error()},nil
		}
		return &userResponse{Msg:"sucess", Err:""},nil
	}
}