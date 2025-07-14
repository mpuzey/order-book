package main

type DepthRequest struct {
	Symbol string `json:"symbol"`
	Limit  *int   `json:"limit"`
}

type DepthResponse struct {
	LastUpdateID int        `json:"lastUpdateId"`
	Bids         [][]string `json:"bids"`
	Asks         [][]string `json:"asks"`
}
