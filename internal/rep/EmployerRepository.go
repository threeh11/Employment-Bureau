package rep

import (
	"errors"
	"gorm.io/gorm"
	"threeh.com/Employment_Bureau/internal/models"
)

type EmployerRepository struct {
	db gorm.DB
}

func NewEmployerRepository(db gorm.DB) *EmployerRepository {
	return &EmployerRepository{
		db: db,
	}
}

func (employerRepository *EmployerRepository) GetById(id string) (*models.Employers, error) {
	if id == "" {
		return nil, errors.New("не передан идентификатор записи")
	}

	employer := &models.Employers{}

	res := employerRepository.db.First(employer, id).Preload("TypesActivity")
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("работодатель не найден")
		}
		return nil, res.Error
	}

	return employer, nil
}

func (employerRepository *EmployerRepository) GetByIds(ids []int) ([]models.Employers, error) {
	if len(ids) == 0 {
		return nil, errors.New("не переданы идентификаторы записи")
	}

	var employers []models.Employers

	res := employerRepository.db.Find(&employers, ids)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("работодатели не найден")
		}
		return nil, res.Error
	}

	return employers, nil
}
