package order

import "docker-go/app/models"

type OrderRepository interface {
	Store(order models.Order) models.Order
	Update(order models.Order) models.Order
	Delete(order models.Order) error
	FindAll(user models.User) []models.Order
	Find(orderId int) models.Order
}
