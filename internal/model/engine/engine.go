package engine

import (
	"sync"
	"trade-engine/internal/enum"
	"trade-engine/internal/model/common"
)

type OrderBook struct {
	sync.RWMutex
	BuyOrders  []common.Order
	SellOrders []common.Order
}

// Add a buy order to the order book
func (book *OrderBook) addBuyOrder(order common.Order) {
	if len(book.BuyOrders) == 0 {
		book.BuyOrders = append(book.BuyOrders, order)
		return
	}

	var count int
	for i, v := range book.BuyOrders {
		count = i
		if v.Price > order.Price {
			break
		}

		if i == len(book.BuyOrders)-1 {
			count++
		}
	}

	if count == len(book.BuyOrders) {
		book.BuyOrders = append(book.BuyOrders, order)
	} else {
		book.BuyOrders = append(book.BuyOrders[:count+1], book.BuyOrders[count:]...)
		book.BuyOrders[count] = order
	}
}

// Add a sell order to the order book
func (book *OrderBook) addSellOrder(order common.Order) {
	if len(book.SellOrders) == 0 {
		book.SellOrders = append(book.SellOrders, order)
		return
	}

	var count int
	for i, v := range book.SellOrders {
		count = i
		if v.Price < order.Price {
			break
		}

		if i == len(book.SellOrders)-1 {
			count++
		}
	}

	if count == len(book.SellOrders) {
		book.SellOrders = append(book.SellOrders, order)
	} else {
		book.SellOrders = append(book.SellOrders[:count+1], book.SellOrders[count:]...)
		book.SellOrders[count] = order
	}
}

// Remove a buy order from the order book at a given index
func (book *OrderBook) removeBuyOrder(index int) {
	book.BuyOrders = append(book.BuyOrders[:index], book.BuyOrders[index+1:]...)
}

// Remove a sell order from the order book at a given index
func (book *OrderBook) removeSellOrder(index int) {
	book.SellOrders = append(book.SellOrders[:index], book.SellOrders[index+1:]...)
}

// Process an order and return the trades generated before adding the remaining amount to the market
func (book *OrderBook) Process(order common.Order) []common.Trade {
	if order.Type == enum.Market {
		if order.Side == enum.Buy {
			return book.processMarketBuy(order)
		}
		return book.processMarketSell(order)
	}
	if order.Side == enum.Buy {
		return book.processLimitBuy(order)
	}
	return book.processLimitSell(order)
}

// Process a market buy order
func (book *OrderBook) processMarketBuy(order common.Order) []common.Trade {
	book.Lock()
	defer book.Unlock()

	trades := make([]common.Trade, 0, 1)
	n := len(book.SellOrders)
	// check if we have at least one matching order
	if n != 0 {
		// traverse all orders that match
		for i := n - 1; i >= 0; i-- {
			// fill the match order
			trade := book.fillOrder(&book.SellOrders[i], &order)
			trades = append(trades, trade)

			if book.SellOrders[i].Quantity == 0 {
				book.removeSellOrder(i)
			}

			if order.Quantity == 0 {
				return trades
			}
		}
	}
	// finally add the remaining order to the list
	book.addBuyOrder(order)
	return trades
}

// Process a market sell order
func (book *OrderBook) processMarketSell(order common.Order) []common.Trade {
	book.Lock()
	defer book.Unlock()

	trades := make([]common.Trade, 0, 1)
	n := len(book.BuyOrders)
	// check if we have at least one matching order
	if n != 0 {
		// traverse all orders that match
		for i := n - 1; i >= 0; i-- {
			// fill the match order
			trade := book.fillOrder(&book.BuyOrders[i], &order)
			trades = append(trades, trade)

			if book.BuyOrders[i].Quantity == 0 {
				book.removeBuyOrder(i)
			}

			if order.Quantity == 0 {
				return trades
			}
		}
	}
	// finally add the remaining order to the list
	book.addSellOrder(order)
	return trades
}

// Process a limit buy order
func (book *OrderBook) processLimitBuy(order common.Order) []common.Trade {
	book.Lock()
	defer book.Unlock()

	trades := make([]common.Trade, 0, 1)
	n := len(book.SellOrders)
	// check if we have at least one matching order
	if n != 0 || book.SellOrders[n-1].Price <= order.Price {
		// traverse all orders that match
		for i := n - 1; i >= 0; i-- {
			if book.SellOrders[i].Price > order.Price {
				break
			}
			// fill the match order
			trade := book.fillOrder(&book.SellOrders[i], &order)
			trades = append(trades, trade)

			if book.SellOrders[i].Quantity == 0 {
				book.removeSellOrder(i)
			}

			if order.Quantity == 0 {
				return trades
			}
		}
	}
	// finally add the remaining order to the list
	book.addBuyOrder(order)
	return trades
}

// Process a limit sell order
func (book *OrderBook) processLimitSell(order common.Order) []common.Trade {
	book.Lock()
	defer book.Unlock()

	trades := make([]common.Trade, 0, 1)
	n := len(book.BuyOrders)
	// check if we have at least one matching order
	if n != 0 || book.BuyOrders[n-1].Price >= order.Price {
		// traverse all orders that match
		for i := n - 1; i >= 0; i-- {
			if book.BuyOrders[i].Price < order.Price {
				break
			}
			// fill the match order
			trade := book.fillOrder(&book.BuyOrders[i], &order)
			trades = append(trades, trade)

			if book.BuyOrders[i].Quantity == 0 {
				book.removeBuyOrder(i)
			}

			if order.Quantity == 0 {
				return trades
			}
		}
	}
	// finally add the remaining order to the list
	book.addSellOrder(order)
	return trades
}

// Process fill match order
func (book *OrderBook) fillOrder(makerOrder, takerOrder *common.Order) common.Trade {
	trade := common.Trade{
		TakerOrdID: takerOrder.ReqOrdID,
		MakerOrdID: makerOrder.ReqOrdID,
		Price:      makerOrder.Price,
	}
	// fill the entire order
	if makerOrder.Quantity >= takerOrder.Quantity {
		trade.Quantity = takerOrder.Quantity
		makerOrder.Quantity -= takerOrder.Quantity
		takerOrder.Quantity = 0
	}
	// fill a partial order and continue
	if makerOrder.Quantity < takerOrder.Quantity {
		trade.Quantity = makerOrder.Quantity
		takerOrder.Quantity -= makerOrder.Quantity
		makerOrder.Quantity = 0
	}

	return trade
}
