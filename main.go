package main

import (
	"net/http"
)

var orderBook = NewOrderBook()

type DepthResponse struct {
	LastUpdateID int        `json:"lastUpdateId"`
	Bids         [][]string `json:"bids"`
	Asks         [][]string `json:"asks"`
}

func main() {
	http.HandleFunc("/api/v3/depth/", DepthRequestHandler)
	http.ListenAndServe(":3001", nil)
}
