package infrastucture

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SqlHandler struct {
	db *gorm.DB
}

func NewSqlHandler() database.SqlHandler {
	dsn := "host=localhost user=threeh password=asd890jkl dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.db = db
	return sqlHandler
}

func (handler *SqlHandler) Create(obj interface{}) {
	handler.db.Create(obj)
}

func (handler *SqlHandler) FindAll(obj interface{}) {
	handler.db.Find(obj)
}

func (handler *SqlHandler) DeleteById(obj interface{}, id string) {
	handler.db.Delete(obj, id)
}

db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})