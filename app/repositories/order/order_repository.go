package order

import (
	"docker-go/app/models"

	"gorm.io/gorm"
)

type orderConnection struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderConnection{
		db: db,
	}
}

func (repo *orderConnection) Store(order models.Order) models.Order {
	repo.db.Create(&order)
	return order
}

func (repo *orderConnection) Update(order models.Order) models.Order {
	repo.db.Model(&models.Order{}).Where("id=?", order.ID).Updates(order)
	return order
}

func (repo *orderConnection) Delete(order models.Order) error {
	repo.db.Delete(&order)
	return nil
}

func (repo *orderConnection) FindAll(user models.User) []models.Order {
	var orders []models.Order
	repo.db.Where("user_id=?", user.ID).Find(&orders)
	return orders
}

func (repo *orderConnection) Find(orderId int) models.Order {
	var order models.Order
	repo.db.Find(&order, orderId)
	return order
}
