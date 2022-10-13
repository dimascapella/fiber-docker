package order

import (
	"docker-go/app/models"
	"docker-go/app/requests"
)

type OrderService interface {
	Store(order requests.OrderCreateRequest, user models.User) models.Order
	Update(order requests.OrderUpdateRequest) models.Order
	Delete(order models.Order) error
	FindAll(user models.User) []models.Order
	Find(orderId int) models.Order
}
