package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetUpDb() *gorm.DB {
	dsn := "host= user= password= dbname= port= sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
