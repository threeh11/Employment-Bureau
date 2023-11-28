package containers

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"threeh.com/Employment_Bureau/internal/controllers"
)

type MainContainer struct {
	e  *echo.Echo
	db gorm.DB
}

func NewMainContainer(e *echo.Echo, db gorm.DB) *MainContainer {
	return &MainContainer{
		e:  e,
		db: db,
	}
}

func (mainContainer *MainContainer) Build() {
	mainController := controllers.NewMainController()
	mainContainer.e.GET("/", func(c echo.Context) error {
		return mainController.Index(c)
	})
}
