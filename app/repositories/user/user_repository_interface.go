package user

import "docker-go/app/models"

type UserRepository interface {
	Store(user models.User) models.User
	Update(user models.User) models.User
	Delete(user models.User) error
	Find(userId int) models.User
	FindAll() []models.User
}
