package models

type Vacancy struct {
	ID          int    `gorm:"primaryKey"`
	NameVacancy string `gorm:"column:name_vacancy"`
	IDActivity  int    `gorm:"column:id_activity"`
	Description string `gorm:"column:description"`
	GetMoney    int    `gorm:"column:get_money"`
}
