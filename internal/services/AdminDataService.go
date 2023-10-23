package services

import (
	"errors"
	"threeh.com/Employment_Bureau/internal/models"
	"threeh.com/Employment_Bureau/internal/rep"
)

type AdminDataService struct {
	employerRep   *rep.EmployerRepository
	dealingRep    *rep.DealingRepository
	jobSeekersRep *rep.JobSeekersRepository
}

func NewAdminDataService(
	employerRep *rep.EmployerRepository,
	dealingRep *rep.DealingRepository,
	jobSeekersRep *rep.JobSeekersRepository,
) *AdminDataService {
	return &AdminDataService{
		employerRep:   employerRep,
		dealingRep:    dealingRep,
		jobSeekersRep: jobSeekersRep,
	}
}

type DataDealingDto struct {
	ID         int
	JobSeekers models.JobSeekers
	Employer   models.Employers
	Post       string
	Commission int
}

func (adminDataService *AdminDataService) GetAllDataForDealing(dealings []models.Dealing) ([]DataDealingDto, error) {
	if len(dealings) == 0 {
		return nil, errors.New("отсутсвуют записи")
	}

	var idsJobSeekers []int
	var idsEmployers []int

	for _, dealing := range dealings {
		idsJobSeekers = append(idsJobSeekers, dealing.IDJobSeekers)
		idsEmployers = append(idsEmployers, dealing.IDEmployer)
	}

	if len(idsJobSeekers) == 0 {
		return nil, errors.New("отсутсвуют id-шники соискателей")
	} else if len(idsEmployers) == 0 {
		return nil, errors.New("отсутсвуют id-шники работодателей")
	}

	employersByIds, err := adminDataService.employerRep.GetByIds(idsEmployers)
	if err != nil {
		return nil, errors.New("не удалось получить работодателей по заданным id")
	}
	jobSeekersByIds, err := adminDataService.jobSeekersRep.GetByIds(idsJobSeekers)
	if err != nil {
		return nil, errors.New("не удалось получить соискателей по заданным id")
	}

	jobSeekersMap := make(map[int]models.JobSeekers)
	for _, seeker := range jobSeekersByIds {
		jobSeekersMap[seeker.ID] = seeker
	}
	employerMap := make(map[int]models.Employers)
	for _, employer := range employersByIds {
		employerMap[employer.ID] = employer
	}

	var viewDataDtos []DataDealingDto
	for _, dealing := range dealings {
		viewDataDtos = append(viewDataDtos, DataDealingDto{
			ID:         dealing.ID,
			JobSeekers: jobSeekersMap[dealing.IDJobSeekers],
			Employer:   employerMap[dealing.IDEmployer],
			Post:       dealing.Post,
			Commission: dealing.Commission,
		})
	}
	return viewDataDtos, nil
}
