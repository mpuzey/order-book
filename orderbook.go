package main

import (
	"sort"
	"strconv"
	"sync"
)

type OrderBook struct {
	lastUpdateID int
	Bids         map[int]Order
	Asks         map[int]Order
	mu           sync.RWMutex
}

func NewOrderBook() *OrderBook {
	return &OrderBook{
		lastUpdateID: 1027024,
		Bids:         make(map[int]Order),
		Asks:         make(map[int]Order),
	}
}

func (ob *OrderBook) AddOrder(order Order) {
	ob.mu.Lock()
	defer ob.mu.Unlock()

	if order.Side == Buy {
		ob.Bids[order.ID] = order
	} else {
		ob.Asks[order.ID] = order
	}

	ob.lastUpdateID++
}

func (ob *OrderBook) CancelOrder(orderID int) {
	ob.mu.Lock()
	defer ob.mu.Lock()
	delete(ob.Bids, orderID)
	delete(ob.Asks, orderID)
}

func (ob *OrderBook) GetDepth(limit int) *DepthResponse {
	// ob.mu.Lock()
	// defer ob.mu.Unlock()

	return &DepthResponse{
		LastUpdateID: ob.lastUpdateID,
		Asks:         ob.GetAsks(limit),
		Bids:         ob.getBids(limit),
	}
}

func (ob *OrderBook) getBids(limit int) [][2]string {
	bidDepth := make(map[float64]float64)

	for _, order := range ob.Bids {
		bidDepth[order.Price] += order.Quantity
	}

	bidPrices := make([]float64, 0, len(bidDepth))
	for price := range bidDepth {
		bidPrices = append(bidPrices, price)
	}

	sort.Sort(sort.Reverse(sort.Float64Slice(bidPrices)))

	bids := make([][2]string, 0, limit)
	for _, price := range bidPrices {
		bids = append(bids, [2]string{
			strconv.FormatFloat(price, 'f', 8, 64),
			strconv.FormatFloat(bidDepth[price], 'f', 8, 64),
		})
		if len(bids) >= limit {
			break
		}
	}
	return bids
}

func (ob *OrderBook) GetAsks(limit int) [][2]string {
	askDepth := make(map[float64]float64)

	for _, order := range ob.Asks {
		askDepth[order.Price] += order.Quantity
	}

	askPrices := make([]float64, 0, len(askDepth))
	for price := range askDepth {
		askPrices = append(askPrices, price)
	}

	sort.Sort(sort.Reverse(sort.Float64Slice(askPrices)))

	asks := make([][2]string, 0, limit)
	for _, price := range askPrices {
		asks = append(asks, [2]string{
			strconv.FormatFloat(price, 'f', 8, 64),
			strconv.FormatFloat(askDepth[price], 'f', 8, 64),
		})
		if len(asks) >= limit {
			break
		}
	}
	return asks
}
