package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/txuna/fileupload/route"
)

var port string = ":1234"

//TemplateRender struct is html template
type TemplateRender struct {
	templates *template.Template
}

//Render func do rendering template
func (t *TemplateRender) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
func main() {
	e := echo.New()
	t := &TemplateRender{
		templates: template.Must(template.ParseGlob("template/*.html")),
	}
	e.Renderer = t
	//Set Static file path
	e.Use(middleware.Static("static"))
	//e.Use(middleware.Logger())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	/* Route*/
	e.GET("/home", route.Home)
	e.GET("/stats", route.Stats)
	e.GET("/statsLog/:page", route.StatsLog)
	e.GET("/homeDirectory/:path", route.HomeDirectory)
	e.GET("/homeDetail/:path", route.HomeDetail)

	e.Logger.Fatal(e.Start(port))
}
