package handlers

import (
	"docker-go/app/models"
	"docker-go/app/requests"
	"docker-go/app/services/user"
	"docker-go/helpers"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	FindAll(ctx *fiber.Ctx) error
	Find(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
}

type userHandler struct {
	service user.UserService
}

func NewUserHandler(handler user.UserService) UserHandler {
	return &userHandler{
		service: handler,
	}
}

func (handler *userHandler) FindAll(ctx *fiber.Ctx) error {
	var users []models.User
	var err error
	users, err = handler.service.FindAll()
	if err != nil {
		errorResponse := helpers.Response{Code: 404, Message: "Failed Find The Data", Data: helpers.EmptyObj{}}
		return ctx.JSON(errorResponse)
	}
	result := helpers.Response{Code: 200, Message: "Success Fetch User", Data: users}
	return ctx.JSON(result)
}

func (handler *userHandler) Find(ctx *fiber.Ctx) error {
	id, _ := strconv.ParseInt(ctx.Params("id"), 10, 32)
	user, err := handler.service.Find(int(id))
	if err != nil || user.ID == 0 {
		errorResponse := helpers.Response{Code: 404, Message: "Failed Find The Data", Data: helpers.EmptyObj{}}
		return ctx.JSON(errorResponse)
	}

	result := helpers.Response{Code: 200, Message: "Success Fetch User", Data: user}
	return ctx.JSON(result)
}

func (handler *userHandler) Create(ctx *fiber.Ctx) error {
	var userCreate requests.UserCreateRequest
	request := ctx.BodyParser(&userCreate)
	if request != nil {
		errorResponse := helpers.Response{Code: 404, Message: "Fill The Form Please", Data: helpers.EmptyObj{}}
		return ctx.JSON(errorResponse)
	}

	user, err := handler.service.Store(userCreate)
	if err != nil {
		errorResponse := helpers.Response{Code: 400, Message: "Error Create Data", Data: helpers.EmptyObj{}}
		return ctx.JSON(errorResponse)
	}
	result := helpers.Response{Code: 200, Message: "Success Create User", Data: user}
	return ctx.JSON(result)
}
