package models

type OrderType string

const (
    Buy  OrderType = "buy"
    Sell OrderType = "sell"
)

type Order struct {
    ID        string    `json:"id"`
    UserID    string    `json:"user_id"`
    Type      OrderType `json:"type"`
    Price     float64   `json:"price"`
    Quantity  float64   `json:"quantity"`
    Timestamp int64     `json:"timestamp"`
}
