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

func (eR *EmployerRepository) GetById(id string) (*models.Employers, error) {
	if id == "" {
		return nil, errors.New("не передан идентификатор записи")
	}

	employer := &models.Employers{}

	res := eR.db.First(employer, id)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("работодатель не найден")
		}
		return nil, res.Error
	}

	return employer, nil
}

//func (er *EmployerRepository) createRow() (int, erorr) {
//}
