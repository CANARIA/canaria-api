package db

import (
	"github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/CANARIA/canaria-api/env"
)

func Init() *gorm.DB {

	conf := env.GetDBConfig()

	db, err := gorm.Open("mysql", conf.Master.GetMySQLDataSource())

	if err != nil {
		logrus.Error(err)
	}

	if conf.LogMode {
		db.LogMode(true)
	}
	db.DB().SetMaxOpenConns(conf.Master.MaxConnections)
	db.DB().SetMaxIdleConns(conf.Master.MaxIdleConnections)


	return db

}
