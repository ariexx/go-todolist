package controller

import (
	"github.com/gofiber/fiber/v2"
	"go-todolist/helper"
	"go-todolist/request"
	"go-todolist/services"
)

type authControllerImpl struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) AuthController {
	return &authControllerImpl{authService}
}

func (a *authControllerImpl) Login(c *fiber.Ctx) error {
	var input request.LoginRequest
	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFail(
			"Bad request",
			"error",
			err.Error()))
	}
	token, err := a.authService.Login(input)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(helper.ApiResponseFail(
			"Unauthorized",
			"error",
			err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(helper.ApiResponseSuccess(fiber.StatusOK, "success", "success", map[string]string{
		"token": token,
		"type":  "Bearer",
	}))
}

func (a *authControllerImpl) Register(c *fiber.Ctx) error {
	var input request.RegisterRequest
	err := c.BodyParser(&input)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFail(
			"Bad request",
			"error",
			err.Error()))
	}
	token, err := a.authService.Register(input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFail(
			"Bad request",
			"error",
			err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(helper.ApiResponseSuccess(fiber.StatusOK, "success", "success", map[string]string{
		"token": token,
	}))
}
