package handlers

import "docker-go/app/services/order"

type orderHandler struct {
	handler order.OrderService
}

func NewOrderHandler(handler order.OrderService) *orderHandler {
	return &orderHandler{
		handler: handler,
	}
}
