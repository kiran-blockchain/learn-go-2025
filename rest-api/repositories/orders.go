
package repositories

import("rest-api/models"
"rest-api/interfaces")

type orderRepository struct {
    orders []models.Order
}

func NewOrderRepository() interfaces.OrderRepository {
    return &orderRepository{
        orders: []models.Order{},
    }
}

func (r *orderRepository) AddOrder(order models.Order) {
    r.orders = append(r.orders, order)
}

func (r *orderRepository) GetOrderBook() []models.Order {
    return r.orders
}

func (r *orderRepository) RemoveOrder(orderID string) {
    for i, o := range r.orders {
        if o.ID == orderID {
            r.orders = append(r.orders[:i], r.orders[i+1:]...)
            return
        }
    }
}

func (r *orderRepository) UpdateOrder(updatedOrder models.Order) {
    for i, o := range r.orders {
        if o.ID == updatedOrder.ID {
            r.orders[i] = updatedOrder
            return
        }
    }
}