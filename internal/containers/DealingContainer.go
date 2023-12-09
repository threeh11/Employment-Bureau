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

func (dCon *DealingContainer) Build() {
	dealingRep := rep.NewDealingRepository(dCon.db)
	typesRep := rep.NewTypesActivitiesRepository(dCon.db)
	dealingController := controllers.NewDealingController(dealingRep, typesRep)
	dCon.e.GET("/vacancy", func(c echo.Context) error {
		return dealingController.Index(c)
	})
	dCon.e.GET("/create-vacancy", func(c echo.Context) error {
		return dealingController.CrateVacancyIndex(c)
	})
	dCon.e.POST("/create-vacancy/save", func(c echo.Context) error {
		return dealingController.SaveVacancy(c)
	})
}
