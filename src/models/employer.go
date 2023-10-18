package models

type employer struct {
	ID              uint   `json:"id" gorm:"primary_key"`
	nameEmployers   string `name:"name_employer"`
	idActivity      int    `name:"id_activity" gorm:"index"`
	addressEmployer string `name:"address_employer"`
	numberEmployer  string `name:"number_employer"`
}
