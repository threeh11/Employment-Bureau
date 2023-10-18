package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetUpDb() *gorm.DB {
	dsn := "host=localhost user=threeh password=asd890jkl dbname=employment_bureau port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
