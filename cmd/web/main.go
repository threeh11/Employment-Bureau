package main

import (
	"github.com/labstack/echo/v4"
	"threeh.com/Employment_Bureau/internal/app"
	"threeh.com/Employment_Bureau/internal/database/postgres"
)

func main() {
	e := echo.New()
	db := postgres.SetUpDb()
	app.RunApp(*db, e)
}
