package helper

import (
	"github.com/spf13/viper"
	"time"
)

func SetConfig(key string, defaultVal string) string {
	viper.SetConfigType("json")
	viper.AddConfigPath("./")
	viper.SetConfigName("app.config")

	viper.SetDefault(key, defaultVal)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	return viper.GetString(key)
}

func Now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
