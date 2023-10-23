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
	adminDataService *services.AdminDataService
}

func NewAdminController(
	dealingRep *rep.DealingRepository,
	adminDataService *services.AdminDataService,
) *AdminController {
	return &AdminController{
		dealingRep:       dealingRep,
		adminDataService: adminDataService,
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
		http.Error(c.Response().Writer, fmt.Sprintf("Ошибка в при получения значения из базы:", err.Error()), http.StatusInternalServerError)
	}

	dataView, err := adminController.adminDataService.GetAllDataForDealing(dealings)
	if err != nil {
		http.Error(c.Response().Writer, fmt.Sprintf("Ошибка в при получения значения из базы:", err.Error()), http.StatusInternalServerError)
	}

	t := template.Must(template.ParseFiles("./../../resources/views/admin/dealing.html"))
	return t.Execute(c.Response(), dataView)
}
