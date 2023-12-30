package main

import (
	"example/httpserver/config"
	"example/httpserver/server"
	"fmt"
	"log"
)

func main() {
	app := new(server.App)
	app.Config = config.Initconfig()
	app.Router = server.InitServer()
	app.InitRoutes()
	if err := app.Router.Run(fmt.Sprintf("%s:%v", app.Config.Server.Host, app.Config.Server.Port)); err != nil {
		log.Default().Printf("[Error] [Server Start] : %s", err)
		panic("coudn't start the server")
	}
}
