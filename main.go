package main

import (
	"app/config"
	"app/routes"
)

func main() {

	config.ConnectDB()

	e := routes.Init()

	e.Logger.Fatal(e.Start("127.0.0.1:3090"))
}
