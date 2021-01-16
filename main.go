package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/txuna/fileupload/route"
)

var port string = "1234"

func main() {
	e := echo.New()
	//Set Static file path
	e.Use(middleware.Static("static"))

	/* Route*/
	e.GET("/", route.Hello)
	e.GET("/home", route.Home)
	e.GET("/stats", route.Stats)

	e.Logger.Fatal(e.Start(port))
}
