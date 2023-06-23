package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DbConn = InitDatabase()

func InitDatabase() *gorm.DB {
	dbHost := DbHost
	dbName := DbName
	dbUser := DbUser
	dbPass := DbPass
	dbPort := DbPort

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		NewLoggerService().SaveLog(logrus.ErrorLevel, err)
		panic(err)
	}

	return db
}

func CloseConnection(db *gorm.DB) {
	//close connection database
	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}

	err = sqlDb.Close()
	if err != nil {
		panic(err)
	}
}
