package main

type Order struct {
	ID        int       `json:"id"`
	Price     float64   `json:"price"`
	Quantity  float64   `json:"quantity"`
	Side      OrderSide `json:"side"`
	Timestamp int64     `json:"timestamp"`
}

type OrderSide string

const (
	Buy  OrderSide = "buy"
	Sell OrderSide = "sell"
)
