package models

type Vacancy struct {
	ID          int `gorm:"primaryKey"`
	NameVacancy string
	IDActivity  int
	Activity    TypesActivity `gorm:"foreignKey:IDActivity"`
	Description string
	GetMoney    int
}
