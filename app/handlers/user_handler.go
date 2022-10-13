package handlers

import "docker-go/app/services/user"

type userHandler struct {
	handler user.UserService
}

func NewUserHandler(handler user.UserService) *userHandler {
	return &userHandler{
		handler: handler,
	}
}
