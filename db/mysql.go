package db

import (
	"github.com/CANARIA/canaria-api/config"
	"github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func Init() *gorm.DB {

	session := getSession()

	return session

}

func getSession() *gorm.DB {

	db, err := gorm.Open("mysql",
		config.USER+":"+config.PASSWORD+"@tcp("+config.HOST+":"+config.PORT+")/"+config.DB+"?parseTime=true")

	if err != nil {
		logrus.Error(err)
	} else {
		db.LogMode(true)
		db.DB().SetMaxOpenConns(5)
		db.DB().SetMaxIdleConns(5)
	}
	return db
}
