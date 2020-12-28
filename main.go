package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Parker-Yang/srv-braveTroops/global"
	"github.com/Parker-Yang/srv-braveTroops/routes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var monitoredSignals = []os.Signal{
	syscall.SIGHUP,
	syscall.SIGINT,
	syscall.SIGTERM,
	syscall.SIGQUIT,
}

func main() {
	logrus.Println(global.Conf.Client)
	gin.SetMode(global.Conf.App.Mode)
	engine := gin.Default()
	routes.SetRoutes(engine)
	go func() {
		err := engine.Run(global.Conf.App.Port)
		logrus.Fatalf("failed to run project! %v", err)
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, monitoredSignals...)
	select {
	case <-quit:
		logrus.Println("exit...")
	default:
	}
}
