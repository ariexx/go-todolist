package config

import "go-todolist/helper"

var (
	AppName = helper.SetConfig("appName", "")
	AppPort = helper.SetConfig("server.port", "")
	DbUser  = helper.SetConfig("database.username", "")
	DbPass  = helper.SetConfig("database.password", "")
	DbHost  = helper.SetConfig("database.host", "")
	DbPort  = helper.SetConfig("database.port", "")
	DbName  = helper.SetConfig("database.database", "")
)
