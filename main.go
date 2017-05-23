package main

import (
	"demo/config"
	"github.com/BurntSushi/toml"
	"demo/log"
	"flag"
	"demo/user"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

var filePath = flag.String("config", "app.toml", "configuration file path")

func main() {
	flag.Parse()

	//loading configurations
	var conf config.ApplicationConfig
	if _, err := toml.DecodeFile(*filePath, &conf); err != nil {
		log.Log("err", err.Error())
		return
	}

	var repo = user.NewRepository(conf.Mysql)
	var svc = user.NewService(repo)
	fmt.Println(svc.Register("test1","test"))

}