package controllers

import (
	"github.com/labstack/echo/v4"
	"html/template"
)

type MainController struct {
}

func NewMainController() *MainController {
	return &MainController{}
}

func (mainController *MainController) Index(c echo.Context) error {
	t := template.Must(template.ParseFiles("./../../resources/views/main/index.html"))
	return t.Execute(c.Response().Writer, nil)
}
