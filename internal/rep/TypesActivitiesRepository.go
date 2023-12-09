package rep

import (
	"errors"
	"gorm.io/gorm"
	"threeh.com/Employment_Bureau/internal/models"
)

type TypesActivitiesRepository struct {
	db gorm.DB
}

func NewTypesActivitiesRepository(db gorm.DB) *TypesActivitiesRepository {
	return &TypesActivitiesRepository{
		db: db,
	}
}

func (tAR *TypesActivitiesRepository) GetAllActivities() ([]models.TypesActivity, error) {
	var typesActivity []models.TypesActivity
	result := tAR.db.Find(&typesActivity)
	if result.Error != nil {
		return nil, errors.New("не найдено не одной записи")
	}
	return typesActivity, nil
}
