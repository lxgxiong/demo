package main

import (
	"demo/config"
	"github.com/BurntSushi/toml"
	"demo/log"
	"flag"
	"demo/user"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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
	var svc user.Service
	svc = user.NewService(repo)
	svc = user.NewLoggingService(log.Logger, svc)

	muxS := http.NewServeMux()
	muxS.Handle("/", user.MakeHandler(svc))
	http.Handle("/", accessControl(muxS))
	errs := make(chan error, 2)
	go func() {
		log.Log("transport", "http", "address", fmt.Sprint(":", conf.Port), "msg", "listening")
		errs <- http.ListenAndServe(fmt.Sprint(":", conf.Port), nil)
	}()
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	log.Log("terminated", <-errs)

}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}