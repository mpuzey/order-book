package main

import (
	"sort"
	"strconv"
	"sync"
)

type OrderBook struct {
	lastUpdateID int
	Bids         map[float64]float64
	Asks         map[float64]float64
	mu           sync.RWMutex
}

func NewOrderBook() *OrderBook {
	return &OrderBook{
		lastUpdateID: 1027024,
		Bids:         make(map[float64]float64),
		Asks:         make(map[float64]float64),
	}
}

func (ob *OrderBook) AddOrder(order Order) {
	ob.mu.Lock()
	defer ob.mu.Unlock()

	if order.Side == Buy {
		ob.addBuy(order)
	} else {
		ob.addSell(order)
	}

}

func (ob *OrderBook) addSell(order Order) {
	if order.Quantity == 0 {
		delete(ob.Asks, order.Price)
	} else {
		ob.Asks[order.Price] += order.Quantity
	}

	ob.lastUpdateID++
}

func (ob *OrderBook) addBuy(order Order) {
	if order.Quantity == 0 {
		delete(ob.Bids, order.Price)
	} else {
		ob.Bids[order.Price] += order.Quantity
	}

	ob.lastUpdateID++
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
	// Convert and sort Bids (desc)
	bidPrices := make([]float64, 0, len(ob.Bids))
	for price := range ob.Bids {
		bidPrices = append(bidPrices, price)
	}
	sort.Sort(sort.Reverse(sort.Float64Slice(bidPrices)))

	bids := make([][2]string, 0, limit)
	for _, price := range bidPrices {
		bids = append(bids, [2]string{
			strconv.FormatFloat(price, 'f', 8, 64),
			strconv.FormatFloat(ob.Bids[price], 'f', 8, 64),
		})
		if len(bids) >= limit {
			break
		}
	}
	return bids
}

func (ob *OrderBook) GetAsks(limit int) [][2]string {
	// Convert and sort Bids (desc)
	askPrices := make([]float64, 0, len(ob.Asks))
	for price := range ob.Asks {
		askPrices = append(askPrices, price)
	}
	sort.Sort(sort.Reverse(sort.Float64Slice(askPrices)))

	asks := make([][2]string, 0, limit)
	for _, price := range askPrices {
		asks = append(asks, [2]string{
			strconv.FormatFloat(price, 'f', 8, 64),
			strconv.FormatFloat(ob.Asks[price], 'f', 8, 64),
		})
		if len(asks) >= limit {
			break
		}
	}
	return asks
}
