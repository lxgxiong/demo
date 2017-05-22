package user

import (
	"github.com/go-xorm/xorm"
	"demo/config"
	"log"
)

type repository struct {
	engine *xorm.Engine
}

func NewRepository(conf config.MysqlConfig) Repository  {
	engine,err := xorm.NewEngine("mysql", conf.User+":"+conf.Password+"@tcp("+conf.Host+":"+conf.Port+")/"+conf.Database)
	if err !=nil{
		log.Fatal("error when initiating databse",err)
	}
	engine.Sync2(new(User))
	return &repository{
		engine:engine,
	}
}

func (r *repository)Register(user *User) (int64,error) {
	return r.engine.InsertOne(user)
}

func (r *repository)Login(user *User) (bool,error){
	return r.engine.Where("username=?",user.Username).And("password=?",user.Password).Get(&user)
}

func (r *repository) ChangePassword(user *User)(int64,error){
	return r.engine.Where("username=?",user.Username).Cols("password").Update(&user)
}
func (r *repository) Delete(user *User)(int64,error){
	return r.engine.Where("username=?",user.Username).And("password=?",user.Password).Delete(&user)
}
