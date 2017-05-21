package user

import "time"

type User struct {
	Id int64
	username string	`xorm:"unique"`
	password string
	created time.Time
	updated time.Time
}

type Repository interface {
	Register(username, password string)
	Login(username,password string)
	ChangePassword(username, password string)
	Delete(username string)
}

