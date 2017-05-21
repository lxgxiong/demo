package user

type Service interface {
	Register(username, password string)
	Login(username, password string)
	ChangePassword(username, password string)
	Delete(username string)
}

type service struct {

}

func (s *service)Register(username, password string) {

}

func (s *service)Login(username, password string) {

}

func (s *service)ChangePassword(username, password string) () {

}

func (s *service)Delete(username string) {

}