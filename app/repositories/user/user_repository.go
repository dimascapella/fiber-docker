package user

import (
	"docker-go/app/models"

	"gorm.io/gorm"
)

type userConnection struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		db: db,
	}
}

func (repo *userConnection) Store(user models.User) models.User {
	repo.db.Create(&user)
	return user
}
func (repo *userConnection) Update(user models.User) models.User {
	repo.db.Preload("orders").Model(&models.User{}).Where("id=?", user.ID).Updates(user)
	return user
}
func (repo *userConnection) Delete(user models.User) error {
	repo.db.Delete(&user)
	return nil
}
func (repo *userConnection) Find(userId int) models.User {
	var user models.User
	repo.db.Preload("orders").Find(&user, userId)
	return user
}
func (repo *userConnection) FindAll() []models.User {
	var users []models.User
	repo.db.Preload("orders").Find(&users)
	return users
}
