package user

import (
	"docker-go/app/models"
	"docker-go/app/repositories/user"
	"docker-go/app/requests"

	"github.com/mashingan/smapping"
)

type userService struct {
	repo user.UserRepository
}

func NewUserService(service user.UserRepository) UserService {
	return &userService{
		repo: service,
	}
}

func (service *userService) Store(user requests.UserCreateRequest) (models.User, error) {
	newUser := models.User{}
	err := smapping.FillStruct(&newUser, smapping.MapFields(&user))
	if err != nil {
		panic("Failed Mapping The Field")
	}
	return service.repo.Store(newUser), nil
}

func (service *userService) Update(user requests.UserUpdateRequest) (models.User, error) {
	updateUser := models.User{}
	err := smapping.FillStruct(&updateUser, smapping.MapFields(&user))
	if err != nil {
		panic("Failed Mapping The Field")
	}
	return service.repo.Update(updateUser), nil
}

func (service *userService) Delete(user models.User) error {
	return service.repo.Delete(user)
}

func (service *userService) Find(userId int) (models.User, error) {
	return service.repo.Find(userId), nil
}

func (service *userService) FindAll() ([]models.User, error) {
	return service.repo.FindAll(), nil
}
