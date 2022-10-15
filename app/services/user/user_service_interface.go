package user

import (
	"docker-go/app/models"
	"docker-go/app/requests"
)

type UserService interface {
	Store(user requests.UserCreateRequest) (models.User, error)
	Update(user requests.UserUpdateRequest) (models.User, error)
	Delete(user models.User) error
	Find(userId int) (models.User, error)
	FindAll() ([]models.User, error)
}
