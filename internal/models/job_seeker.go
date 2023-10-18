package models

type JobSeekers struct {
	ID                  int `gorm:"primaryKey"`
	SurnameSeekers      string
	NameSeekers         string
	PatronymicSeekers   string
	QualificationSeeker string
	IDActivity          int
	Activity            TypesActivity `gorm:"foreignKey:IDActivity"`
	OtherData           string
	NeedMoney           int
}
