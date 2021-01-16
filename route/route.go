package route

import (
	"net/http"

	"github.com/labstack/echo"
)

//Hello is test Request function - e.GET
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}

//Home response home.html
func Home(c echo.Context) error {
	return c.File("template/home.html")
}

//Stats response stats.html
func Stats(c echo.Context) error {
	return c.File("template/stats.html")
}
