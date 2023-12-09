package services

import (
	"threeh.com/Employment_Bureau/internal/rep"
)

type VacancyDataService struct {
	dealingRep *rep.DealingRepository
}

func NewVacancyDataService(dealingRep *rep.DealingRepository) *VacancyDataService {
	return &VacancyDataService{
		dealingRep: dealingRep,
	}
}

type VacancyDto struct {
}
