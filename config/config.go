package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func SetConfig(key string, defaultVal string) string {
	viper.SetConfigType("json")
	viper.AddConfigPath("./")
	viper.SetConfigName("app.config")

	viper.SetDefault(key, defaultVal)
	err := viper.ReadInConfig()
	if err != nil {
		NewLoggerService().SaveLog(logrus.ErrorLevel, err)
	}

	return viper.GetString(key)
}

var (
	DbUser = SetConfig("database.user", "root")
	DbPass = SetConfig("database.pass", "")
	DbHost = SetConfig("database.host", "localhost")
	DbPort = SetConfig("database.port", "3306")
	DbName = SetConfig("database.name", "belajar_golang_todolist")
)
