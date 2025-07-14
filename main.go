package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var orderBook = NewOrderBook()

func DepthRequestHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("testing!")
	if r.Method == http.MethodGet {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		return
	}

	var req DepthRequest
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	depthResponse := &DepthResponse{
		LastUpdateID: orderBook.lastUpdateID,
		Bids:         orderBook.GetBids(),
		Asks:         orderBook.GetAsks(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(depthResponse)

}

func main() {
	http.HandleFunc("/api/v3/depth/", DepthRequestHandler)
	http.ListenAndServe(":3001", nil)
}
