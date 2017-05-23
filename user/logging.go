package user

import (
	"github.com/go-kit/kit/log"
	"time"
)

type loggingService struct {
	logger log.Logger
	Service
}

func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}
func (s *loggingService)Register(username, password string) (user *User, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "register",
			"took", time.Since(begin),
			"username", username,
		)
	}(time.Now())
	return s.Service.Register(username, password)
}
func (s *loggingService)Login(username, password string) (bool, error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "login",
			"took", time.Since(begin),
			"username", username,
		)
	}(time.Now())
	return s.Service.Login(username, password)
}
func (s *loggingService)ChangePassword(username, password string) (bool, error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "ChangePassword",
			"took", time.Since(begin),
			"username", username,
		)
	}(time.Now())
	return s.Service.ChangePassword(username, password)
}
func (s *loggingService)Delete(username, password string) (int64, error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "Delete",
			"took", time.Since(begin),
			"username", username,
		)
	}(time.Now())
	return s.Service.Delete(username, password)
}