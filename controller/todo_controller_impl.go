package controller

import (
	"github.com/gofiber/fiber/v2"
	"go-todolist/helper"
	"go-todolist/request"
	"go-todolist/services"
)

type todoControllerImpl struct {
	todoService services.TodoService
}

func (t *todoControllerImpl) GetByUserId(ctx *fiber.Ctx) error {

	//pada bagian ini kita mengambil user_id dari context yang sudah kita set pada middleware
	//karena kita sudah menambahkan middleware JwtMiddlewares() pada route todo
	//maka kita bisa mengambil user_id dari context
	//jika user_id tidak ada pada context, maka akan mengembalikan error
	//dengan ini masalah authorization sudah selesai
	userId := ctx.Locals("user_id").(uint)

	query, err := t.todoService.GetByUserId(userId)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(helper.ApiResponseFail("Todo not found", "error", err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(helper.ApiResponseSuccess(200, "success", "success", query))
}

func (t *todoControllerImpl) Create(ctx *fiber.Ctx) error {
	var input request.CreateTodoRequest
	err := ctx.BodyParser(&input)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFail(
			"Bad request",
			"error",
			err.Error()))
	}

	userId := ctx.Locals("user_id").(uint)
	_, err = t.todoService.Create(userId, &input)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(helper.ApiResponseFail(
			"Failed to create todo",
			"error",
			err.Error()))
	}

	return ctx.Status(fiber.StatusCreated).JSON(helper.ApiResponseSuccess(200, "success", "success", nil))

}

func (t *todoControllerImpl) Update(userId uint, todoId int, todo any) error {
	//TODO implement me
	panic("implement me")
}

func (t *todoControllerImpl) Delete(userId uint, todoId int) error {
	//TODO implement me
	panic("implement me")
}

func NewTodoController(todoService services.TodoService) TodoController {
	return &todoControllerImpl{todoService}
}
