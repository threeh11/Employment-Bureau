package models

type Employers struct {
	ID           int           `gorm:"primaryKey"`
	NameEmployer string        `gorm:"column:name_employer"`
	IDActivity   int           `gorm:"column:id_activity"`
	Activity     TypesActivity `gorm:"foreignKey:IDActivity"`
	Address      string        `gorm:"column:address_employer"`
	Number       string        `gorm:"column:number_employre"`
}
