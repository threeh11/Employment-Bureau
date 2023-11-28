package models

type Dealing struct {
	ID           int    `gorm:"primaryKey"`
	IDJobSeekers int    `gorm:"column:id_job_seekers"`
	IDEmployer   int    `gorm:"column:id_employer"`
	IDActivity   int    `gorm:"column:id_activity"`
	IsVacancy    int    `gorm:"column:is_vacancy"`
	Post         string `gorm:"column:post"`
	Commission   int    `gorm:"column:commission"`
	Description  string `gorm:"column:description"`
	GetMoney     int    `gorm:"column:get_money"`
}
