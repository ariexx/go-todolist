package controller

import (
	"github.com/gofiber/fiber/v2"
	"go-todolist/helper"
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
		return ctx.Status(fiber.StatusNotFound).JSON(helper.ApiResponseFail(
			"User not found",
			"error",
			err.Error()))
	}
	return ctx.Status(fiber.StatusOK).JSON(helper.ApiResponseSuccess(
		fiber.StatusOK,
		"success",
		"success",
		user))
}

func (u *userControllerImpl) CreateUser(ctx *fiber.Ctx) error {
	var input request.CreateUserRequest
	err := ctx.BodyParser(&input)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFail(
			"Bad request",
			"error",
			err.Error()))
	}
	_, err = u.service.CreateUser(input)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(helper.ApiResponseFail(
			"Failed to create user",
			"error",
			err.Error()))
	}
	return ctx.Status(fiber.StatusCreated).JSON(helper.ApiResponseSuccess(
		fiber.StatusCreated,
		"success",
		"success",
		[]string{}))
}
