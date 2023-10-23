package rep

import (
	"errors"
	"gorm.io/gorm"
	"threeh.com/Employment_Bureau/internal/models"
)

type DealingRepository struct {
	db gorm.DB
}

func NewDealingRepository(db gorm.DB) *DealingRepository {
	return &DealingRepository{
		db: db,
	}
}

func (dR *DealingRepository) GetAllDataForAdmin() ([]models.Dealing, error) {
	var dealing []models.Dealing
	result := dR.db.Find(&dealing)
	if result.Error != nil {
		return nil, errors.New("не найдено не одной записи")
	}
	return dealing, nil
}
