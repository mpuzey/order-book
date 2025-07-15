package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID        int       `json:"id"`
	Price     float64   `json:"price"`
	Quantity  float64   `json:"quantity"`
	Side      OrderSide `json:"side"`
	Timestamp int64     `json:"timestamp"`
}

type OrderSide string

const (
	Buy  OrderSide = "BUY"
	Sell OrderSide = "SELL"
)

type OrderResponse struct {
	OrderID int `json:"orderId"`
}

func OrderRequestHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	}

	var req Order
	body := json.NewDecoder(r.Body)
	if err := body.Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}

	id, _ := strconv.Atoi(uuid.New().String())
	req.ID = id
	req.Timestamp = time.Now().UnixMilli()

	orderBook.AddOrder(req)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&OrderResponse{
		OrderID: req.ID,
	})

}
