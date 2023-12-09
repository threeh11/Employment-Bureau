package controllers

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"html/template"
	"net/http"
	"threeh.com/Employment_Bureau/internal/rep"
)

type DealingController struct {
	dealingRep *rep.DealingRepository
	typesRep   *rep.TypesActivitiesRepository
}

func NewDealingController(
	dealingRep *rep.DealingRepository,
	typesRep *rep.TypesActivitiesRepository,
) *DealingController {
	return &DealingController{
		dealingRep: dealingRep,
		typesRep:   typesRep,
	}
}

func (dC *DealingController) Index(c echo.Context) error {
	t := template.Must(template.ParseFiles("./../../resources/views/vacancy/index.html"))
	return t.Execute(c.Response().Writer, nil)
}

func (dC *DealingController) CrateVacancyIndex(c echo.Context) error {
	types, err := dC.typesRep.GetAllActivities()
	if err != nil {
		mes := fmt.Sprintf("Ошибка при получения значения из базы:", err.Error())
		http.Error(c.Response().Writer, mes, http.StatusInternalServerError)
	}
	t := template.Must(template.ParseFiles("./../../resources/views/vacancy/index.html"))
	return t.Execute(c.Response().Writer, types)
}

// Vacancy Валидация
type Vacancy struct {
	Post          string `json:"post" form:"post" validate:"required,max=200,excludesall=0123456789"`
	Commission    int    `json:"commission" form:"commission" validate:"required,min=1,max=200"`
	Description   string `json:"description" form:"description" validate:"required,max=1000,excludesall=0123456789"`
	GetMoney      int    `json:"get_money" form:"get_money" validate:"required,min=1,max=1000000"`
	TypesActivity int    `json:"id_activity" form:"id_activity" validate:"required,min=1,max=1000000"`
}

func (v *Vacancy) validate() error {
	validate := validator.New()
	return validate.Struct(v)
}

func (dC *DealingController) SaveVacancy(c echo.Context) error {
	vacancy := new(Vacancy)
	if err := c.Bind(vacancy); err != nil {
		return c.HTML(http.StatusBadRequest, "Неверный формат данных")
	}

	// Валидация полей
	if err := vacancy.validate(); err != nil {
		// Обработка ошибки валидации
		mes := fmt.Sprintf("Ошибка валидации полей:", err.Error())
		return c.HTML(http.StatusBadRequest, mes)
	}

	//err := dC.dealingRep.SaveVacancy(vacancy)

	// Возврат успешного ответа
	return c.HTML(http.StatusOK, "Вакансия успешно сохранена!")
}
