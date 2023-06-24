package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"go-todolist/config"
	"go-todolist/controller"
	"go-todolist/repository"
	"go-todolist/services"
)

func main() {
	db := config.DbConn
	defer config.CloseConnection(db)

	app := fiber.New()

	// Routes
	userRepository := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	router := app.Group("/api/v1")
	router.Post("/users", userController.CreateUser)
	router.Get("/users/:id", userController.GetUserById)

	err := app.Listen(":" + config.AppPort)
	if err != nil {
		config.NewLoggerService().SaveLog(logrus.ErrorLevel, err)
		panic(err)
	}
}
