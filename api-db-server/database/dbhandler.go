package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type RDBHandler struct {
	UserName      string
	Password      string
	ServerAddress string
	DbName        string
	db            *gorm.DB
}

func (rdb *RDBHandler) Connect() error {
	dsn := rdb.UserName + ":" + rdb.Password + "@tcp(" + rdb.ServerAddress + ")/" + rdb.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	log.Println("database::dsn:", dsn)
	var err error
	rdb.db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("database::Connect > Error:", err)
		return err
	}

	return nil
}
