package models

type Dealing struct {
	ID           int    `gorm:"primaryKey"`
	IDJobSeekers int    `gorm:"column:id_job_seekers"`
	IDEmployer   int    `gorm:"column:id_employer"`
	Post         string `gorm:"column:post"`
	Commission   int    `gorm:"column:commission"`
}
