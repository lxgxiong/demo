package user

import "time"

type User struct {
	Id       int64
	Username string `xorm:"unique"`
	Password string
	Created  time.Time `xorm:"created"`
	Updated  time.Time `xorm:"updated"`
}

type Repository interface {
	Register(user *User) (int64, error)
	Login(user *User) (bool, error)
	ChangePassword(user *User) (int64, error)
	Delete(user *User)(int64,error)
}
