package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"go-todolist/config"
)

func main() {
	db := config.DbConn
	defer config.CloseConnection(db)

	app := fiber.New()

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.Send([]byte("Anjayy"))
	})

	err := app.Listen(":" + config.AppPort)
	if err != nil {
		config.NewLoggerService().SaveLog(logrus.ErrorLevel, err)
		panic(err)
	}
}
