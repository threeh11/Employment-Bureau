package containers

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"threeh.com/Employment_Bureau/internal/controllers"
	"threeh.com/Employment_Bureau/internal/rep"
)

type AdminContainer struct {
	e  *echo.Echo
	db gorm.DB
}

func NewAdminContainer(e *echo.Echo, db gorm.DB) *AdminContainer {
	return &AdminContainer{
		e:  e,
		db: db,
	}
}

func (aC *AdminContainer) Build() {
	employerRep := rep.NewEmployerRepository(aC.db)
	adminController := controllers.NewAdminController(*employerRep)
	aC.e.GET("/admin", func(c echo.Context) error {
		return adminController.Index(c.Response(), c.Request())
	})
}
