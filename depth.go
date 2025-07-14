package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type DepthResponse struct {
	LastUpdateID int        `json:"lastUpdateId"`
	Bids         [][]string `json:"bids"`
	Asks         [][]string `json:"asks"`
}

func DepthRequestHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		return
	}

	q, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		panic(err)
	}
	fmt.Println("symbol is:", q.Get("symbol"))

	depthResponse := &DepthResponse{
		LastUpdateID: orderBook.lastUpdateID,
		Bids:         orderBook.GetBids(),
		Asks:         orderBook.GetAsks(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(depthResponse)

}
