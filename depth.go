package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type DepthResponse struct {
	LastUpdateID int         `json:"lastUpdateId"`
	Bids         [][2]string `json:"bids"`
	Asks         [][2]string `json:"asks"`
}

func DepthRequestHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("symbol is:", r.URL.Query().Get("symbol"))

	limitStr := r.URL.Query().Get("limit")
	limit := 100
	if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
		limit = l
	}

	depthResponse := orderBook.GetDepth(limit)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(depthResponse)

}
