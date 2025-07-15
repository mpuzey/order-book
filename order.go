package main

import (
	"encoding/json"
	"fmt"
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
	Buy OrderSide = "BUY"
)

type OrderResponse struct {
	OrderID int `json:"orderId"`
}

type Order2 struct {
	ID        int    `json:"id"`
	Price     string `json:"price"`
	Quantity  string `json:"quantity"`
	Side      string `json:"side"`
	Timestamp int    `json:"timestamp"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var req Order2
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Invalid JSON: %+v", err), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Got order: %+v", req)
}

func OrderRequestHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	}

	// body, _ := io.ReadAll(r.Body)
	// fmt.Println("Raw body:", string(body))

	decoder := json.NewDecoder(r.Body)
	var req Order
	if err := decoder.Decode(&req); err != nil {
		errStr := fmt.Sprintf("Invalid JSON: %+v", err)
		http.Error(w, errStr, http.StatusBadRequest)
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
