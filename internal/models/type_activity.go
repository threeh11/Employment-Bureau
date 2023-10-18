package models

type TypesActivity struct {
	ID           int `gorm:"primaryKey"`
	NameActivity string
}
