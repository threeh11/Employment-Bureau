package models

type TypesActivity struct {
	ID           int    `gorm:"primaryKey"`
	NameActivity string `gorm:"column:name_activity"`
}
