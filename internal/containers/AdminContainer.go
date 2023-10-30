package containers

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"threeh.com/Employment_Bureau/internal/controllers"
	"threeh.com/Employment_Bureau/internal/rep"
	"threeh.com/Employment_Bureau/internal/services"
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

func (adminContainer *AdminContainer) Build() {

	// -- Admin
	dealingRep := rep.NewDealingRepository(adminContainer.db)
	employerRep := rep.NewEmployerRepository(adminContainer.db)
	jobSeekersRep := rep.NewJobSeekersRepository(adminContainer.db)
	adminDataService := services.NewAdminDataService(employerRep, dealingRep, jobSeekersRep)
	adminController := controllers.NewAdminController(dealingRep, employerRep, jobSeekersRep, adminDataService)
	adminContainer.e.GET("/admin", func(c echo.Context) error {
		return adminController.Index(c)
	})
	adminContainer.e.GET("/admin/dealing", func(c echo.Context) error {
		return adminController.Dealing(c)
	})
	adminContainer.e.GET("/admin/employers", func(c echo.Context) error {
		return adminController.Employers(c)
	})
	adminContainer.e.GET("/admin/job_seekers", func(c echo.Context) error {
		return adminController.JobSeekers(c)
	})
}
