package database

import (
	"fmt"
	"ginweibo/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/lexkong/log"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	db, err := gorm.Open(config.DBConfig.Connection, config.DBConfig.URL)
	if err != nil {
		log.Fatal("Database connection failed. URL: "+config.DBConfig.URL+" error: ", err)
	} else {
		fmt.Print("\n\n------- GORM OPEN SUCCESS! -------\n\n")
	}
	db = db.Set("gorm:table_options", "ENGINE=InnoDB  DEFAULT CHARSET=utf8;").AutoMigrate()
	db.LogMode(config.DBConfig.Debug)
	DB = db
	return db
}
