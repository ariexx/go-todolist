package controller

import (
	"github.com/gofiber/fiber/v2"
	"go-todolist/request"
	"go-todolist/services"
	"strconv"
)

type userControllerImpl struct {
	service services.UserService
}

func NewUserController(service services.UserService) UserController {
	return &userControllerImpl{service: service}
}

func (u *userControllerImpl) GetUserById(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	user, err := u.service.GetUserById(id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
}

func (u *userControllerImpl) CreateUser(ctx *fiber.Ctx) error {
	var input request.CreateUserRequest
	err := ctx.BodyParser(&input)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	_, err = u.service.CreateUser(input)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
	})
}
