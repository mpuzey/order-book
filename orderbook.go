package main

import (
	"sort"
	"strconv"
	"sync"
)

type OrderBook struct {
	lastUpdateID int
	Bids         []Order
	Asks         []Order
	mu           sync.RWMutex
}

func NewOrderBook() *OrderBook {
	return &OrderBook{
		lastUpdateID: 1027024,
		Bids: []Order{
			{
				Price:    4.00000000,
				Quantity: 431.00000000,
				Side:     Buy,
			},
		},
		Asks: []Order{
			{
				Price:    4.00000200,
				Quantity: 12.00000000,
				Side:     Sell,
			},
		},
	}
}

func (ob *OrderBook) AddOrder(order Order) {
	ob.mu.Lock()
	defer ob.mu.Unlock()

	if order.Side == Buy {
		ob.Bids = append(ob.Bids, order)
		sort.SliceStable(ob.Bids, func(i, j int) bool {
			if ob.Bids[i].Price == ob.Bids[j].Price {
				return ob.Bids[i].Timestamp < ob.Bids[j].Timestamp
			}
			return ob.Bids[i].Price > ob.Bids[j].Price
		})
	} else {
		ob.Asks = append(ob.Asks, order)
		sort.SliceStable(ob.Asks, func(i, j int) bool {
			if ob.Asks[i].Price == ob.Asks[j].Price {
				return ob.Asks[i].Timestamp < ob.Asks[j].Timestamp
			}
			return ob.Asks[i].Price < ob.Asks[j].Price
		})
	}

	ob.lastUpdateID = order.ID
}

func (ob *OrderBook) GetBids() [][]string {
	ob.mu.RLock()
	defer ob.mu.RUnlock()

	var bids [][]string
	for _, v := range ob.Bids {

		price := strconv.FormatFloat(v.Price, 'E', -1, 64)
		quantity := strconv.FormatFloat(v.Quantity, 'E', -1, 64)
		bids = append(bids, []string{price, quantity})
	}

	return bids
}

func (ob *OrderBook) GetAsks() [][]string {
	ob.mu.RLock()
	defer ob.mu.RUnlock()

	var asks [][]string
	for _, v := range ob.Asks {
		price := strconv.FormatFloat(v.Price, 'E', -1, 64)
		quantity := strconv.FormatFloat(v.Quantity, 'E', -1, 64)
		asks = append(asks, []string{price, quantity})
	}

	return asks
}
