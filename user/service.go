package user

type Service interface {
	Register(username, password string) (user *User, err error)
	Login(username, password string) (bool, error)
	ChangePassword(username, password string)(bool, error)
	Delete(username string)
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
	return s.repo.Login(user)
}

func (s *service)ChangePassword(username, password string) (bool, error) {
	user := &User{
		Username:username,
		Password:password,
	}
	s.repo.ChangePassword(user)
	return true,nil
}

func (s *service)Delete(username string) {

}