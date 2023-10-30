package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"html/template"
	"net/http"
	"threeh.com/Employment_Bureau/internal/rep"
	"threeh.com/Employment_Bureau/internal/services"
)

type AdminController struct {
	dealingRep       *rep.DealingRepository
	employerRep      *rep.EmployerRepository
	jobSeekerRep     *rep.JobSeekersRepository
	adminDataService *services.AdminDataService
}

func NewAdminController(
	dealingRep *rep.DealingRepository,
	employerRep *rep.EmployerRepository,
	jobSeekerRep *rep.JobSeekersRepository,
	adminDataService *services.AdminDataService,
) *AdminController {
	return &AdminController{
		dealingRep:       dealingRep,
		adminDataService: adminDataService,
		employerRep:      employerRep,
		jobSeekerRep:     jobSeekerRep,
	}
}

func (adminController *AdminController) Index(c echo.Context) error {
	// to do Обработать_!
	t := template.Must(template.ParseFiles("./../../resources/views/admin/index.html"))
	return t.Execute(c.Response().Writer, nil)
}

func (adminController *AdminController) Dealing(c echo.Context) error {
	dealings, err := adminController.dealingRep.GetAllDataForAdmin()
	if err != nil {
		mes := fmt.Sprintf("Ошибка при получения значения из базы:", err.Error())
		http.Error(c.Response().Writer, mes, http.StatusInternalServerError)
	}

	dataView, err := adminController.adminDataService.GetAllDataForDealing(dealings)
	if err != nil {
		mes := fmt.Sprintf("Ошибка при получения значения из базы:", err.Error())
		http.Error(c.Response().Writer, mes, http.StatusInternalServerError)
	}

	t := template.Must(template.ParseFiles("./../../resources/views/admin/dealing.html"))
	return t.Execute(c.Response(), dataView)
}

func (adminController *AdminController) Employers(c echo.Context) error {
	employers, err := adminController.employerRep.GetAllDataForAdmin()
	if err != nil {
		mes := fmt.Sprintf("Ошибка при получения значения из базы:", err.Error())
		http.Error(c.Response().Writer, mes, http.StatusInternalServerError)
	}

	t := template.Must(template.ParseFiles("./../../resources/views/admin/employer.html"))
	return t.Execute(c.Response(), employers)
}

func (adminController *AdminController) JobSeekers(c echo.Context) error {
	jobSeekers, err := adminController.jobSeekerRep.GetAllDataForAdmin()
	if err != nil {
		mes := fmt.Sprintf("Ошибка при получения значения из базы:", err.Error())
		http.Error(c.Response().Writer, mes, http.StatusInternalServerError)
	}

	t := template.Must(template.ParseFiles("./../../resources/views/admin/job_seekers.html"))
	return t.Execute(c.Response(), jobSeekers)
}
