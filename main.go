package main

import (
	"github.com/rk-the-dev/micro-fiber-svc/app"
	"github.com/rk-the-dev/micro-fiber-svc/cmd"
	"github.com/rk-the-dev/micro-fiber-svc/helpers/config"
)

func main() {
	config.ExportENV()
	app.InitHelpers()
	app.Logger.Infoln("Bootstrapping the service...")
	cmd.Run()
}
