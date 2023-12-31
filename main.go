package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"go-todolist/config"
	"go-todolist/controller"
	"go-todolist/middlewares"
	"go-todolist/repository"
	"go-todolist/services"
)

func main() {
	db := config.DbConn
	defer config.CloseConnection(db)

	app := fiber.New()

	// Middlewares
	app.Use(middlewares.GeneralMiddleware())
	app.Use(middlewares.LogMiddleware())

	// Routes
	userRepository := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	todoRepository := repository.NewTodoRepository(db)
	todoService := services.NewTodoService(todoRepository)
	todoController := controller.NewTodoController(todoService)

	jwtService := services.NewJwtService()
	authService := services.NewAuthService(userRepository, jwtService)
	authController := controller.NewAuthController(authService)

	authRouter := app.Group("/api/auth")
	authRouter.Post("/login", authController.Login)
	authRouter.Post("/register", authController.Register)

	router := app.Group("/api/v1", middlewares.JwtMiddlewares(), middlewares.SignatureMiddleware())
	router.Get("/users/:id", userController.GetUserById)
	router.Get("/todos", todoController.GetByUserId)
	router.Post("/todos", todoController.Create)

	err := app.Listen(":" + config.AppPort)
	if err != nil {
		config.NewLoggerService().SaveLog(logrus.ErrorLevel, err)
		panic(err)
	}
}
