package main

import (
	"app/config"
	"app/routes"

	"github.com/labstack/echo/v4/middleware"
)

func main() {

	config.ConnectDB()

	e := routes.Init()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "${time_custom} [ ${status} ] ${remote_ip} ${method} ${uri} in ${latency_human}\n",
		CustomTimeFormat: "[ 15:04:05 | 02-01-2006 ]",
	}))

	e.Logger.Fatal(e.Start("127.0.0.1:3090"))
}
