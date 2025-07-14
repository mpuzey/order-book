package main

import (
	"encoding/json"
	"net/http"
)

type Trade struct {
	ID           int64  `json:"id"`
	Price        string `json:"price"`
	Qty          string `json:"qty"`
	QuoteQty     string `json:"quoteQty"`
	Time         int64  `json:"time"`
	IsBuyerMaker bool   `json:"isBuyerMaker"`
	IsBestMatch  bool   `json:"isBestMatch"`
}

type TradesRequest struct {
	Symbol string `json:"symbol"`
	Limit  *int   `json:"limit"`
}

// Get recent trades
func tradesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}
	var req TradesRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	trades := []Trade{
		{
			ID:           28457,
			Price:        "4.00000100",
			Qty:          "12.00000000",
			QuoteQty:     "48.000012",
			Time:         1499865549590,
			IsBuyerMaker: true,
			IsBestMatch:  true,
		},
		{
			ID:           28458,
			Price:        "4.00000200",
			Qty:          "5.00000000",
			QuoteQty:     "20.000010",
			Time:         1499865550000,
			IsBuyerMaker: false,
			IsBestMatch:  true,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trades)
}
