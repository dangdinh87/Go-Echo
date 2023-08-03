package main

import (
	"github.com/labstack/echo/v4"
	"myapp-echo/configs"
	"myapp-echo/routes"
)

func main() {

	e := echo.New()
	routes.UserRoute(e) //add this

	//run database
	configs.ConnectDB()
	e.Logger.Fatal(e.Start(":6000"))

}
