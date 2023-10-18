package main

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"net/http"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Hello(c echo.Context) error {
	return c.Render(http.StatusOK, "hello", "World")
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Главная страница")
	})
	t := &Template{
		templates: template.Must(template.ParseGlob("resources/views/*.html")),
	}
	e.Renderer = t
	e.GET("/hello", Hello)
	e.Logger.Fatal(e.Start(":1323"))
}
