package main

import (
	"demo/conf"
	"demo/dao"
	"demo/server/http"
	"demo/service"
	"github.com/google/wire"
	"os"
	"os/signal"
	"syscall"
)
func main()  {
	conf.InitConfig()
	wire.Build(conf.InitConfig, dao.New, service.New, http.Init)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			os.Exit(0)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}

}
