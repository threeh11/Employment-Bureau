package models

type jobSeekers struct {
	ID                int    `json:"id" gorm:"primary_key"`
	surnameSeekers    string `name:"surname_seekers"`
	nameSeekers       string `name:"name_seekers" gorm:"index"`
	patronymicSeekers string `name:"patronymic_seekers"`
}
