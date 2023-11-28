package containers

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"threeh.com/Employment_Bureau/internal/controllers"
	"threeh.com/Employment_Bureau/internal/rep"
)

type DealingContainer struct {
	e  *echo.Echo
	db gorm.DB
}

func NewDealingContainer(e *echo.Echo, db gorm.DB) *DealingContainer {
	return &DealingContainer{
		db: db,
		e:  e,
	}
}

func (dealingContainer *DealingContainer) Build() {
	dealingRep := rep.NewDealingRepository(dealingContainer.db)
	dealingController := controllers.NewDealingController(dealingRep)
	dealingContainer.e.GET("/vacancy", func(c echo.Context) error {
		return dealingController.Index(c)
	})
}
