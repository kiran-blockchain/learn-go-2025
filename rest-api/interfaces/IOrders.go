package interfaces

import (
	"rest-api/models"
)

type OrderRepository interface {
	AddOrder(order models.Order)
	GetOrderBook() []models.Order
	RemoveOrder(orderID string)
	UpdateOrder(order models.Order)
}
