package order

import (
	"docker-go/app/models"
	"docker-go/app/repositories/order"
	"docker-go/app/requests"

	"github.com/mashingan/smapping"
)

type orderService struct {
	repo order.OrderRepository
}

func NewOrderService(repo order.OrderRepository) OrderService {
	return &orderService{
		repo: repo,
	}
}

func (service *orderService) Store(order requests.OrderCreateRequest, user models.User) models.Order {
	createOrder := models.Order{}
	err := smapping.FillStruct(&createOrder, smapping.MapFields(&order))
	if err != nil {
		panic("Failed Mapping The Field")
	}
	createOrder.UserID = user.ID
	return service.repo.Store(createOrder)
}

func (service *orderService) Update(order requests.OrderUpdateRequest) models.Order {
	updateOrder := models.Order{}
	err := smapping.FillStruct(&updateOrder, smapping.MapFields(&order))
	if err != nil {
		panic("Failed Mapping The Field")
	}
	return service.repo.Update(updateOrder)
}

func (service *orderService) Delete(order models.Order) error {
	return service.repo.Delete(order)
}

func (service *orderService) FindAll(user models.User) []models.Order {
	return service.repo.FindAll(user)
}

func (service *orderService) Find(orderId int) models.Order {
	return service.repo.Find(orderId)
}
