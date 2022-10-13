package user

import (
	"docker-go/app/models"
	"docker-go/app/requests"
)

type UserService interface {
	Store(user requests.UserCreateRequest) models.User
	Update(user requests.UserUpdateRequest) models.User
	Delete(user models.User) error
	Find(userId int) models.User
	FindAll() []models.User
}
