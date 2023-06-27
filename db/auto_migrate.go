package db

import (
	"github.com/sirupsen/logrus"
	"go-todolist/config"
	"go-todolist/model"
)

// AutoMigrateDb is a function to migrate all models
func AutoMigrateDb() {
	db := config.DbConn
	defer config.CloseConnection(db)
	err := db.AutoMigrate(
		&model.User{})
	if err != nil {
		config.NewLoggerService().SaveLog(logrus.ErrorLevel, err)
		panic(err)
	}
}
