package rep

import (
	"errors"
	"gorm.io/gorm"
	"threeh.com/Employment_Bureau/internal/models"
)

type JobSeekersRepository struct {
	db gorm.DB
}

func NewJobSeekersRepository(db gorm.DB) *JobSeekersRepository {
	return &JobSeekersRepository{
		db: db,
	}
}

func (jobSeekersRepository *JobSeekersRepository) GetByIds(ids []int) ([]models.JobSeekers, error) {
	if len(ids) == 0 {
		return nil, errors.New("не переданы идентификаторы записи")
	}

	var jobSeekers []models.JobSeekers

	res := jobSeekersRepository.db.Find(&jobSeekers, ids)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("работодатели не найден")
		}
		return nil, res.Error
	}

	return jobSeekers, nil
}

func (jobSeekersRepository *JobSeekersRepository) GetAllDataForAdmin() ([]models.JobSeekers, error) {
	var jobSeekers []models.JobSeekers
	result := jobSeekersRepository.db.Find(&jobSeekers)
	if result.Error != nil {
		return nil, errors.New("не найдено не одной записи")
	}
	return jobSeekers, nil
}
