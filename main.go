package main

import (
	"net/http"
)

var orderBook = NewOrderBook()

func main() {
	http.HandleFunc("/api/v3/depth/", DepthRequestHandler)
	http.HandleFunc("/api/v3/order/", OrderRequestHandler)
	http.ListenAndServe(":3001", nil)
}
