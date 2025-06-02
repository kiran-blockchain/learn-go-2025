package services

import (
    "rest-api/models"
     repositories "rest-api/interfaces"
    "sort"
)

type MatchingService interface {
    PlaceOrder(order models.Order) []models.Order
}

type matchingService struct {
    repo repositories.OrderRepository
}

func NewMatchingService(r repositories.OrderRepository) MatchingService {
    return &matchingService{repo: r}
}

func (s *matchingService) PlaceOrder(order models.Order) []models.Order {
    matchedTrades := []models.Order{}
    orderBook := s.repo.GetOrderBook()

    if order.Type == models.Buy {
        // Sort sell orders (ascending price)
        sort.Slice(orderBook, func(i, j int) bool {
            return orderBook[i].Price < orderBook[j].Price
        })
    } else {
        // Sort buy orders (descending price)
        sort.Slice(orderBook, func(i, j int) bool {
            return orderBook[i].Price > orderBook[j].Price
        })
    }

    for i := 0; i < len(orderBook); i++ {
        existingOrder := orderBook[i]
        // Skip if same user (no self-trade)
        if existingOrder.UserID == order.UserID {
            continue
        }

        // Check for matching condition
        if (order.Type == models.Buy && order.Price >= existingOrder.Price) ||
            (order.Type == models.Sell && order.Price <= existingOrder.Price) {

            // Determine trade quantity (partial fills allowed)
            tradeQty := min(order.Quantity, existingOrder.Quantity)

            matchedTrades = append(matchedTrades, models.Order{
                ID:       existingOrder.ID,
                UserID:   existingOrder.UserID,
                Type:     existingOrder.Type,
                Price:    existingOrder.Price,
                Quantity: tradeQty,
            })

            order.Quantity -= tradeQty
            existingOrder.Quantity -= tradeQty

            if existingOrder.Quantity == 0 {
                s.repo.RemoveOrder(existingOrder.ID)
            } else {
                s.repo.UpdateOrder(existingOrder)
            }

            if order.Quantity == 0 {
                break
            }
        }
    }

    // If order has remaining quantity, add to book
    if order.Quantity > 0 {
        s.repo.AddOrder(order)
    }

    return matchedTrades
}

func min(a, b float64) float64 {
    if a < b {
        return a
    }
    return b
}
