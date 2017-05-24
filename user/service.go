package user

import (
	"errors"
)

type Service interface {
	Register(username, password string) (user *User, err error)
	Login(username, password string) (bool, error)
	ChangePassword(username, password,newPass string)(bool, error)
	Delete(username,password string) (int64,error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo:repo,
	}
}

func (s *service)Register(username, password string) (user *User, err error) {
	user = &User{
		Username:username,
		Password:password,
	}
	_, err = s.repo.Register(user)
	if err != nil {
		return nil, err
	}
	return
}

func (s *service)Login(username, password string) (bool, error) {
	user := &User{
		Username:username,
		Password:password,
	}
	b,e := s.repo.Login(user)
	return b,e
}

func (s *service)ChangePassword(username, password,newPass string) (bool, error) {
	effected, err :=s.repo.ChangePassword(username, password,newPass)
	if err!=nil {
		return false,err
	}
	if effected <= 0 {
		return false, errors.New("change password failed, check your username and password.")
	}
	return true,nil
}

func (s *service)Delete(username,password string) (int64,error){
	user := &User{
		Username:username,
		Password:password,
	}
	return s.repo.Delete(user)
}