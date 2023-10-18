package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"threeh.com/Employment_Bureau/internal/rep"
)

type AdminController struct {
	employerRep rep.EmployerRepository
}

func NewAdminController(employerRep rep.EmployerRepository) *AdminController {
	return &AdminController{
		employerRep: employerRep,
	}
}

func (eC *AdminController) Index(w http.ResponseWriter, r *http.Request) error {
	employerID := r.URL.Query().Get("id")

	employer, err := eC.employerRep.GetById(employerID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при получении работодателя: %s", err.Error()), http.StatusInternalServerError)
		return nil
	}

	response, err := json.Marshal(employer)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при конвертации работодателя в JSON: %s", err.Error()), http.StatusInternalServerError)
		return nil
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(response)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка сервера: %s", err.Error()), http.StatusInternalServerError)
		return nil
	}
	return nil
}
