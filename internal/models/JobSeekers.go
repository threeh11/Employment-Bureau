package models

type JobSeekers struct {
	ID                  int    `gorm:"primaryKey"`
	SurnameSeekers      string `gorm:"column:surname_seekers"`
	NameSeekers         string `gorm:"column:name_seekers"`
	PatronymicSeekers   string `gorm:"column:patronymic_seekers"`
	QualificationSeeker string `gorm:"column:qualification_seekers"`
	IDActivity          int    `gorm:"column:id_activity"`
	OtherData           string `gorm:"column:other_data"`
	NeedMoney           int    `gorm:"column:need_money_seekers"`
}
