package app

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"threeh.com/Employment_Bureau/internal/containers"
)

func RunApp(db gorm.DB, e *echo.Echo) {
	adminContainer := containers.NewAdminContainer(e, db)
	adminContainer.Build()
	err := e.Start(":8080")
	if err != nil {
		panic(err)
	}
}
