package models

type TypesActivity struct {
	ID           int    `gorm:"primaryKey"`
	NameActivity string `gorm:"column:name_activity"`
}

func (TypesActivity) TableName() string {
	return "types_activity"
}
