package postgres

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func SetUpDb() *gorm.DB {
	err := godotenv.Load("/../../../.env")
	if err != nil {
		panic(err)
	}
	dbHost := os.Getenv("POSTGRES_HOST")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRESS_PASSWORD")
	dbName := os.Getenv("POSTGRESS_DB_NAME")
	dbPort := os.Getenv("POSTGRESS_PORT")
	dsn := "host=" + dbHost +
		" user=" + dbUser +
		" password=" + dbPassword +
		" dbname=" + dbName +
		" port=" + dbPort + " sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
