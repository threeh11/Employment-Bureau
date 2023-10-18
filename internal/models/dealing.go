package models

type Dealing struct {
	ID           int `gorm:"primaryKey"`
	IDJobSeekers int
	JobSeeker    JobSeekers `gorm:"foreignKey:IDJobSeekers"`
	IDEmployer   int
	Employer     Employers `gorm:"foreignKey:IDEmployer"`
	Post         string
	Commission   int
}
