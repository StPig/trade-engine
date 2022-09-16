package engine

import (
	"fmt"
	"testing"
	"trade-engine/internal/model/common"
)

var (
	// M is mean Market
	// L is mean Limit
	// B is mean Buy
	// S is mean Sell
	MBOrder1 = common.Order{
		Name:     "test1",
		ReqOrdID: 1,
		Price:    10,
		Quantity: 1,
		Type:     0,
		Side:     0,
	}

	MBOrder2 = common.Order{
		Name:     "test2",
		ReqOrdID: 1,
		Price:    20,
		Quantity: 1,
		Type:     0,
		Side:     0,
	}

	MBOrder3 = common.Order{
		Name:     "test3",
		ReqOrdID: 1,
		Price:    30,
		Quantity: 1,
		Type:     0,
		Side:     0,
	}

	MBOrder4 = common.Order{
		Name:     "test4",
		ReqOrdID: 1,
		Price:    40,
		Quantity: 1,
		Type:     0,
		Side:     0,
	}

	MBOrder5 = common.Order{
		Name:     "test5",
		ReqOrdID: 1,
		Price:    50,
		Quantity: 1,
		Type:     0,
		Side:     0,
	}

	MBOrder6 = common.Order{
		Name:     "test6",
		ReqOrdID: 1,
		Price:    60,
		Quantity: 1,
		Type:     0,
		Side:     0,
	}

	MBOrder7 = common.Order{
		Name:     "test7",
		ReqOrdID: 1,
		Price:    70,
		Quantity: 1,
		Type:     0,
		Side:     0,
	}

	MSOrder1 = common.Order{
		Name:     "test1",
		ReqOrdID: 1,
		Price:    10,
		Quantity: 1,
		Type:     0,
		Side:     1,
	}

	MSOrder2 = common.Order{
		Name:     "test2",
		ReqOrdID: 1,
		Price:    20,
		Quantity: 1,
		Type:     0,
		Side:     1,
	}

	MSOrder3 = common.Order{
		Name:     "test3",
		ReqOrdID: 1,
		Price:    30,
		Quantity: 1,
		Type:     0,
		Side:     1,
	}

	MSOrder4 = common.Order{
		Name:     "test4",
		ReqOrdID: 1,
		Price:    40,
		Quantity: 1,
		Type:     0,
		Side:     1,
	}

	LBOrder1 = common.Order{
		Name:     "test1",
		ReqOrdID: 1,
		Price:    10,
		Quantity: 1,
		Type:     1,
		Side:     0,
	}

	LBOrder2 = common.Order{
		Name:     "test2",
		ReqOrdID: 1,
		Price:    20,
		Quantity: 1,
		Type:     1,
		Side:     0,
	}

	LSOrder1 = common.Order{
		Name:     "test1",
		ReqOrdID: 1,
		Price:    10,
		Quantity: 1,
		Type:     1,
		Side:     1,
	}

	LSOrder6 = common.Order{
		Name:     "test6",
		ReqOrdID: 1,
		Price:    60,
		Quantity: 1,
		Type:     1,
		Side:     1,
	}
)

func TestAddBuyOrderToEmptySlice(t *testing.T) {
	book := OrderBook{}

	book.addBuyOrder(MBOrder1)

	if len(book.BuyOrders) != 1 {
		t.Fatal("order not insert")
	} else if book.BuyOrders[0] != MBOrder1 {
		t.Fatal("order is not same")
	}
}

func TestAddBuyOrderToLast(t *testing.T) {
	book := OrderBook{
		BuyOrders: []common.Order{MBOrder1, MBOrder2},
	}

	book.addBuyOrder(MBOrder3)

	if len(book.BuyOrders) != 3 {
		t.Fatal("order not insert")
	} else if book.BuyOrders[2] != MBOrder3 {
		t.Fatal("order is not same")
	}
}

func TestAddBuyOrderToFront(t *testing.T) {
	book := OrderBook{
		BuyOrders: []common.Order{MBOrder2, MBOrder3},
	}

	book.addBuyOrder(MBOrder1)

	if len(book.BuyOrders) != 3 {
		t.Fatal("order not insert")
	} else if book.BuyOrders[0] != MBOrder1 {
		t.Fatal("order is not same")
	}
}

func TestAddBuyOrderToMiddle1(t *testing.T) {
	book := OrderBook{
		BuyOrders: []common.Order{MBOrder1, MBOrder3},
	}

	book.addBuyOrder(MBOrder2)

	if len(book.BuyOrders) != 3 {
		t.Fatal("order not insert")
	} else if book.BuyOrders[1] != MBOrder2 {
		t.Fatal("order is not same")
	}
}

func TestAddBuyOrderToMiddle2(t *testing.T) {
	book := OrderBook{
		BuyOrders: []common.Order{MBOrder1, MBOrder2, MBOrder4},
	}

	book.addBuyOrder(MBOrder3)

	if len(book.BuyOrders) != 4 {
		t.Fatal("order not insert")
	} else if book.BuyOrders[2] != MBOrder3 {
		t.Fatal("order is not same")
	}
}

func TestAddSellOrderToEmptySlice(t *testing.T) {
	book := OrderBook{}

	book.addSellOrder(MSOrder1)

	if len(book.SellOrders) != 1 {
		t.Fatal("order not insert")
	} else if book.SellOrders[0] != MSOrder1 {
		t.Fatal("order is not same")
	}
}

func TestAddSellOrderToLast(t *testing.T) {
	book := OrderBook{
		SellOrders: []common.Order{MSOrder3, MSOrder2},
	}

	book.addSellOrder(MSOrder1)

	if len(book.SellOrders) != 3 {
		t.Fatal("order not insert")
	} else if book.SellOrders[2] != MSOrder1 {
		t.Fatal("order is not same")
	}
}

func TestAddSellOrderToFront(t *testing.T) {
	book := OrderBook{
		SellOrders: []common.Order{MSOrder2, MSOrder1},
	}

	book.addSellOrder(MSOrder3)

	if len(book.SellOrders) != 3 {
		t.Fatal("order not insert")
	} else if book.SellOrders[0] != MSOrder3 {
		t.Fatal("order is not same")
	}
}

func TestAddSellOrderToMiddle1(t *testing.T) {
	book := OrderBook{
		SellOrders: []common.Order{MSOrder3, MSOrder1},
	}

	book.addSellOrder(MSOrder2)

	if len(book.SellOrders) != 3 {
		t.Fatal("order not insert")
	} else if book.SellOrders[1] != MSOrder2 {
		t.Fatal("order is not same")
	}
}

func TestAddSellOrderToMiddle2(t *testing.T) {
	book := OrderBook{
		SellOrders: []common.Order{MSOrder4, MSOrder2, MSOrder1},
	}

	book.addSellOrder(MSOrder3)

	if len(book.SellOrders) != 4 {
		t.Fatal("order not insert")
	} else if book.SellOrders[1] != MSOrder3 {
		t.Fatal("order is not same")
	}
}

func TestRemoveBuyOrderFront(t *testing.T) {
	book := OrderBook{
		BuyOrders: []common.Order{MBOrder1, MBOrder2, MBOrder3},
	}

	book.removeBuyOrder(0)

	if len(book.BuyOrders) != 2 {
		t.Fatal("order not remove")
	} else if book.BuyOrders[0] != MBOrder2 {
		t.Fatal("order 0 is not same")
	} else if book.BuyOrders[1] != MBOrder3 {
		t.Fatal("order 1 is not same")
	}
}

func TestRemoveBuyOrderLast(t *testing.T) {
	book := OrderBook{
		BuyOrders: []common.Order{MBOrder1, MBOrder2, MBOrder4},
	}

	book.removeBuyOrder(2)

	if len(book.BuyOrders) != 2 {
		t.Fatal("order not remove")
	} else if book.BuyOrders[0] != MBOrder1 {
		t.Fatal("order 0 is not same")
	} else if book.BuyOrders[1] != MBOrder2 {
		t.Fatal("order 1 is not same")
	}
}

func TestRemoveBuyOrderMiddle1(t *testing.T) {
	book := OrderBook{
		BuyOrders: []common.Order{MBOrder1, MBOrder2, MBOrder4},
	}

	book.removeBuyOrder(1)

	if len(book.BuyOrders) != 2 {
		t.Fatal("order not remove")
	} else if book.BuyOrders[0] != MBOrder1 {
		t.Fatal("order 0 is not same")
	} else if book.BuyOrders[1] != MBOrder4 {
		t.Fatal("order 1 is not same")
	}
}

func TestRemoveBuyOrderMiddle2(t *testing.T) {
	book := OrderBook{
		BuyOrders: []common.Order{MBOrder1, MBOrder2, MBOrder3, MBOrder4},
	}

	book.removeBuyOrder(2)

	if len(book.BuyOrders) != 3 {
		t.Fatal("order not remove")
	} else if book.BuyOrders[0] != MBOrder1 {
		t.Fatal("order 0 is not same")
	} else if book.BuyOrders[1] != MBOrder2 {
		t.Fatal("order 1 is not same")
	} else if book.BuyOrders[2] != MBOrder4 {
		t.Fatal("order 1 is not same")
	}
}

func TestRemoveSellOrderFront(t *testing.T) {
	book := OrderBook{
		SellOrders: []common.Order{MSOrder3, MSOrder2, MSOrder1},
	}

	book.removeSellOrder(0)

	if len(book.SellOrders) != 2 {
		t.Fatal("order not insert")
	} else if book.SellOrders[0] != MSOrder2 {
		t.Fatal("order 0 is not same")
	} else if book.SellOrders[1] != MSOrder1 {
		t.Fatal("order 1 is not same")
	}
}

func TestRemoveSellOrderLast(t *testing.T) {
	book := OrderBook{
		SellOrders: []common.Order{MSOrder4, MSOrder3, MSOrder2},
	}

	book.removeSellOrder(2)

	if len(book.SellOrders) != 2 {
		t.Fatal("order not insert")
	} else if book.SellOrders[0] != MSOrder4 {
		t.Fatal("order 0 is not same")
	} else if book.SellOrders[1] != MSOrder3 {
		t.Fatal("order 1 is not same")
	}
}

func TestRemoveSellOrderMiddle1(t *testing.T) {
	book := OrderBook{
		SellOrders: []common.Order{MSOrder4, MSOrder3, MSOrder1},
	}

	book.removeSellOrder(1)

	if len(book.SellOrders) != 2 {
		t.Fatal("order not insert")
	} else if book.SellOrders[0] != MSOrder4 {
		t.Fatal("order 0 is not same")
	} else if book.SellOrders[1] != MSOrder1 {
		t.Fatal("order 1 is not same")
	}
}

func TestRemoveSellOrderMiddle2(t *testing.T) {
	book := OrderBook{
		SellOrders: []common.Order{MSOrder4, MSOrder3, MSOrder2, MSOrder1},
	}

	book.removeSellOrder(2)

	if len(book.SellOrders) != 3 {
		t.Fatal("order not insert")
	} else if book.SellOrders[0] != MSOrder4 {
		t.Fatal("order 0 is not same")
	} else if book.SellOrders[1] != MSOrder3 {
		t.Fatal("order 1 is not same")
	} else if book.SellOrders[2] != MSOrder1 {
		t.Fatal("order 2 is not same")
	}
}

func TestProcessMarketBuy(t *testing.T) {
	book := OrderBook{
		SellOrders: []common.Order{MSOrder1},
	}

	fakeTrade := common.Trade{
		TakerOrdID: 1,
		MakerOrdID: 1,
		Price:      10,
		Quantity:   1,
	}
	trades := book.Process(MBOrder1)

	if len(book.SellOrders) != 0 {
		t.Fatal("order not insert")
	} else if len(book.BuyOrders) != 0 {
		t.Fatal("invalid buy order length")
	} else if len(trades) != 1 {
		t.Fatal("invalid trade length")
	} else if trades[0] != fakeTrade {
		t.Fatal("invalid trade")
	}
}

func TestProcessMarketSell(t *testing.T) {
	book := OrderBook{
		BuyOrders: []common.Order{MBOrder1},
	}

	fakeTrade := common.Trade{
		TakerOrdID: 1,
		MakerOrdID: 1,
		Price:      10,
		Quantity:   1,
	}
	trades := book.Process(MSOrder1)

	if len(book.BuyOrders) != 0 {
		t.Fatal("order not insert")
	} else if len(book.SellOrders) != 0 {
		t.Fatal("invalid buy order length")
	} else if len(trades) != 1 {
		t.Fatal("invalid trade length")
	} else if trades[0] != fakeTrade {
		t.Fatal("invalid trade")
	}
}

func TestProcessLimitBuy(t *testing.T) {
	book := OrderBook{
		SellOrders: []common.Order{MSOrder1},
	}

	fakeTrade := common.Trade{
		TakerOrdID: 1,
		MakerOrdID: 1,
		Price:      10,
		Quantity:   1,
	}
	trades := book.Process(LBOrder1)

	if len(book.SellOrders) != 0 {
		t.Fatal("order not insert")
	} else if len(book.BuyOrders) != 0 {
		t.Fatal("invalid buy order length")
	} else if len(trades) != 1 {
		t.Fatal("invalid trade length")
	} else if trades[0] != fakeTrade {
		t.Fatal("invalid trade")
	}
}

func TestProcessLimitSell(t *testing.T) {
	book := OrderBook{
		BuyOrders: []common.Order{MBOrder1},
	}

	fakeTrade := common.Trade{
		TakerOrdID: 1,
		MakerOrdID: 1,
		Price:      10,
		Quantity:   1,
	}
	trades := book.Process(LSOrder1)

	if len(book.BuyOrders) != 0 {
		t.Fatal("order not insert")
	} else if len(book.SellOrders) != 0 {
		t.Fatal("invalid buy order length")
	} else if len(trades) != 1 {
		t.Fatal("invalid trade length")
	} else if trades[0] != fakeTrade {
		t.Fatal("invalid trade")
	}
}

func TestProcessMarketBuyNeedRemove(t *testing.T) {
	book := OrderBook{
		SellOrders: []common.Order{MSOrder4, MSOrder3, MSOrder2, MSOrder1},
	}

	fakeTrade := common.Trade{
		TakerOrdID: 1,
		MakerOrdID: 1,
		Price:      10,
		Quantity:   1,
	}
	trades := book.processMarketBuy(MBOrder2)

	if len(book.SellOrders) != 3 {
		t.Fatal("order not process")
	} else if book.SellOrders[0] != MSOrder4 {
		t.Fatal("order 0 is not same")
	} else if book.SellOrders[1] != MSOrder3 {
		t.Fatal("order 1 is not same")
	} else if book.SellOrders[2] != MSOrder2 {
		t.Fatal("order 2 is not same")
	} else if len(trades) != 1 {
		t.Fatal("invalid trade length")
	} else if trades[0] != fakeTrade {
		t.Fatal("invalid trade")
	}
}

func TestProcessMarketBuyAndSellOrderRemaining(t *testing.T) {
	book := OrderBook{
		SellOrders: []common.Order{MSOrder4, MSOrder3, MSOrder2, {
			Name:     "test1",
			ReqOrdID: 1,
			Price:    10,
			Quantity: 2,
			Type:     0,
			Side:     1,
		}},
	}

	fakeTrade := common.Trade{
		TakerOrdID: 1,
		MakerOrdID: 1,
		Price:      10,
		Quantity:   1,
	}
	trades := book.processMarketBuy(MBOrder2)

	if len(book.SellOrders) != 4 {
		t.Fatal("order not process")
	} else if book.SellOrders[0] != MSOrder4 {
		t.Fatal("order 0 is not same")
	} else if book.SellOrders[1] != MSOrder3 {
		t.Fatal("order 1 is not same")
	} else if book.SellOrders[2] != MSOrder2 {
		t.Fatal("order 2 is not same")
	} else if book.SellOrders[3] != MSOrder1 {
		t.Fatal("order 3 is not same")
	} else if len(trades) != 1 {
		t.Fatal("invalid trade length")
	} else if trades[0] != fakeTrade {
		t.Fatal("invalid trade")
	}
}

func TestProcessMarketBuyAndBuyOrderRemaining(t *testing.T) {
	book := OrderBook{
		SellOrders: []common.Order{MSOrder4, MSOrder3, MSOrder2, MSOrder1},
	}

	trade1, trade2, trade3, trade4 := common.Trade{
		TakerOrdID: 1,
		MakerOrdID: 1,
		Price:      10,
		Quantity:   1,
	}, common.Trade{
		TakerOrdID: 1,
		MakerOrdID: 1,
		Price:      20,
		Quantity:   1,
	}, common.Trade{
		TakerOrdID: 1,
		MakerOrdID: 1,
		Price:      30,
		Quantity:   1,
	}, common.Trade{
		TakerOrdID: 1,
		MakerOrdID: 1,
		Price:      40,
		Quantity:   1,
	}

	trades := book.processMarketBuy(common.Order{
		Name:     "test1",
		ReqOrdID: 1,
		Price:    40,
		Quantity: 5,
		Type:     0,
		Side:     0,
	})

	if len(book.SellOrders) != 0 {
		t.Fatal("order not process")
	} else if len(trades) != 4 {
		t.Fatal("invalid trade length")
	} else if trades[0] != trade1 {
		t.Fatal("trade 1 not same")
	} else if trades[1] != trade2 {
		t.Fatal("trade 2 not same")
	} else if trades[2] != trade3 {
		t.Fatal("trade 3 not same")
	} else if trades[3] != trade4 {
		t.Fatal("trade 4 not same")
	} else if len(book.BuyOrders) != 1 {
		t.Fatal("invalid buy order length")
	} else if book.BuyOrders[0] == MBOrder1 {
		t.Fatal("buy order 1 not same")
	}
}

func TestProcessMarketSellNeedRemove(t *testing.T) {
	book := OrderBook{
		BuyOrders: []common.Order{MBOrder1, MBOrder2, MBOrder3, MBOrder4},
	}

	fakeTrade := common.Trade{
		TakerOrdID: 1,
		MakerOrdID: 1,
		Price:      40,
		Quantity:   1,
	}
	trades := book.processMarketSell(MSOrder3)

	if len(book.BuyOrders) != 3 {
		t.Fatal("order not process")
	} else if book.BuyOrders[0] != MBOrder1 {
		t.Fatal("order 0 is not same")
	} else if book.BuyOrders[1] != MBOrder2 {
		t.Fatal("order 1 is not same")
	} else if book.BuyOrders[2] != MBOrder3 {
		t.Fatal("order 2 is not same")
	} else if len(trades) != 1 {
		t.Fatal("invalid trade length")
	} else if trades[0] != fakeTrade {
		t.Fatal("invalid trade")
	}
}

func TestProcessMarketSellAndBuyOrderRemaining(t *testing.T) {
	book := OrderBook{
		BuyOrders: []common.Order{MBOrder1, MBOrder2, MBOrder3, {
			Name:     "test4",
			ReqOrdID: 1,
			Price:    40,
			Quantity: 2,
			Type:     0,
			Side:     0,
		}},
	}

	fakeTrade := common.Trade{
		TakerOrdID: 1,
		MakerOrdID: 1,
		Price:      40,
		Quantity:   1,
	}
	trades := book.processMarketSell(MSOrder3)

	if len(book.BuyOrders) != 4 {
		t.Fatal("order not process")
	} else if book.BuyOrders[0] != MBOrder1 {
		t.Fatal("order 0 is not same")
	} else if book.BuyOrders[1] != MBOrder2 {
		t.Fatal("order 1 is not same")
	} else if book.BuyOrders[2] != MBOrder3 {
		t.Fatal("order 2 is not same")
	} else if book.BuyOrders[3] != MBOrder4 {
		t.Fatal("order 3 is not same")
	} else if len(trades) != 1 {
		t.Fatal("invalid trade length")
	} else if trades[0] != fakeTrade {
		t.Fatal("invalid trade")
	}
}

func TestProcessMarketSellAndSellOrderRemaining(t *testing.T) {
	book := OrderBook{
		BuyOrders: []common.Order{MBOrder1, MBOrder2, MBOrder3, MBOrder4},
	}

	trade1, trade2, trade3, trade4 := common.Trade{
		TakerOrdID: 1,
		MakerOrdID: 1,
		Price:      40,
		Quantity:   1,
	}, common.Trade{
		TakerOrdID: 1,
		MakerOrdID: 1,
		Price:      30,
		Quantity:   1,
	}, common.Trade{
		TakerOrdID: 1,
		MakerOrdID: 1,
		Price:      20,
		Quantity:   1,
	}, common.Trade{
		TakerOrdID: 1,
		MakerOrdID: 1,
		Price:      10,
		Quantity:   1,
	}

	trades := book.processMarketSell(common.Order{
		Name:     "test1",
		ReqOrdID: 1,
		Price:    40,
		Quantity: 5,
		Type:     0,
		Side:     0,
	})

	if len(book.BuyOrders) != 0 {
		t.Fatal("order not process")
	} else if len(trades) != 4 {
		t.Fatal("invalid trade length")
	} else if trades[0] != trade1 {
		t.Fatal("trade 1 not same")
	} else if trades[1] != trade2 {
		t.Fatal("trade 2 not same")
	} else if trades[2] != trade3 {
		t.Fatal("trade 3 not same")
	} else if trades[3] != trade4 {
		t.Fatal("trade 4 not same")
	} else if len(book.SellOrders) != 1 {
		t.Fatal("invalid buy order length")
	} else if book.SellOrders[0] == MSOrder1 {
		t.Fatal("buy order 1 not same")
	}
}

func TestProcessLimitBuyNeedRemove(t *testing.T) {
	book := OrderBook{
		SellOrders: []common.Order{MSOrder4, MSOrder3, MSOrder2, MSOrder1},
	}

	fakeTrade := common.Trade{
		TakerOrdID: 1,
		MakerOrdID: 1,
		Price:      10,
		Quantity:   1,
	}
	trades := book.processLimitBuy(LBOrder2)

	if len(book.SellOrders) != 3 {
		t.Fatal("order not process")
	} else if book.SellOrders[0] != MSOrder4 {
		t.Fatal("order 0 is not same")
	} else if book.SellOrders[1] != MSOrder3 {
		t.Fatal("order 1 is not same")
	} else if book.SellOrders[2] != MSOrder2 {
		t.Fatal("order 2 is not same")
	} else if len(trades) != 1 {
		t.Fatal("invalid trade length")
	} else if trades[0] != fakeTrade {
		t.Fatal("invalid trade")
	}
}

func TestProcessLimitBuyAndSellOrderRemaining(t *testing.T) {
	book := OrderBook{
		SellOrders: []common.Order{MSOrder4, MSOrder3, MSOrder2, {
			Name:     "test1",
			ReqOrdID: 1,
			Price:    10,
			Quantity: 2,
			Type:     0,
			Side:     1,
		}},
	}

	fakeTrade := common.Trade{
		TakerOrdID: 1,
		MakerOrdID: 1,
		Price:      10,
		Quantity:   1,
	}
	trades := book.processLimitBuy(LBOrder2)

	if len(book.SellOrders) != 4 {
		t.Fatal("order not process")
	} else if book.SellOrders[0] != MSOrder4 {
		t.Fatal("order 0 is not same")
	} else if book.SellOrders[1] != MSOrder3 {
		t.Fatal("order 1 is not same")
	} else if book.SellOrders[2] != MSOrder2 {
		t.Fatal("order 2 is not same")
	} else if book.SellOrders[3] != MSOrder1 {
		t.Fatal("order 3 is not same")
	} else if len(trades) != 1 {
		t.Fatal("invalid trade length")
	} else if trades[0] != fakeTrade {
		t.Fatal("invalid trade")
	}
}

func TestProcessLimitBuyAndBuyOrderRemaining(t *testing.T) {
	book := OrderBook{
		SellOrders: []common.Order{MSOrder4, MSOrder3, MSOrder2, MSOrder1},
	}

	trade1, trade2, trade3, trade4 := common.Trade{
		TakerOrdID: 1,
		MakerOrdID: 1,
		Price:      10,
		Quantity:   1,
	}, common.Trade{
		TakerOrdID: 1,
		MakerOrdID: 1,
		Price:      20,
		Quantity:   1,
	}, common.Trade{
		TakerOrdID: 1,
		MakerOrdID: 1,
		Price:      30,
		Quantity:   1,
	}, common.Trade{
		TakerOrdID: 1,
		MakerOrdID: 1,
		Price:      40,
		Quantity:   1,
	}

	remainingOrder := common.Order{
		Name:     "test1",
		ReqOrdID: 1,
		Price:    40,
		Quantity: 1,
		Type:     1,
		Side:     0,
	}

	trades := book.processLimitBuy(common.Order{
		Name:     "test1",
		ReqOrdID: 1,
		Price:    40,
		Quantity: 5,
		Type:     1,
		Side:     0,
	})

	fmt.Println(book.BuyOrders[0])
	if len(book.SellOrders) != 0 {
		t.Fatal("order not process")
	} else if len(trades) != 4 {
		t.Fatal("invalid trade length")
	} else if trades[0] != trade1 {
		t.Fatal("trade 1 not same")
	} else if trades[1] != trade2 {
		t.Fatal("trade 2 not same")
	} else if trades[2] != trade3 {
		t.Fatal("trade 3 not same")
	} else if trades[3] != trade4 {
		t.Fatal("trade 4 not same")
	} else if len(book.BuyOrders) != 1 {
		t.Fatal("invalid buy order length")
	} else if book.BuyOrders[0] != remainingOrder {
		t.Fatal("buy order 1 not same")
	}
}

func TestProcessLimitBuyToLimit(t *testing.T) {
	book := OrderBook{
		SellOrders: []common.Order{MSOrder4, MSOrder3, MSOrder2, MSOrder1},
	}

	trade1, trade2 := common.Trade{
		TakerOrdID: 1,
		MakerOrdID: 1,
		Price:      10,
		Quantity:   1,
	}, common.Trade{
		TakerOrdID: 1,
		MakerOrdID: 1,
		Price:      20,
		Quantity:   1,
	}

	remainingOrder := common.Order{
		Name:     "test1",
		ReqOrdID: 1,
		Price:    20,
		Quantity: 2,
		Type:     1,
		Side:     0,
	}

	trades := book.processLimitBuy(common.Order{
		Name:     "test1",
		ReqOrdID: 1,
		Price:    20,
		Quantity: 4,
		Type:     1,
		Side:     0,
	})

	if len(book.SellOrders) != 2 {
		t.Fatal("order not process")
	} else if len(trades) != 2 {
		t.Fatal("invalid trade length")
	} else if trades[0] != trade1 {
		t.Fatal("trade 1 not same")
	} else if trades[1] != trade2 {
		t.Fatal("trade 2 not same")
	} else if len(book.BuyOrders) != 1 {
		t.Fatal("invalid buy order length")
	} else if book.BuyOrders[0] != remainingOrder {
		t.Fatal("buy order 1 not same")
	}
}

func TestProcessLimitSellNeedRemove(t *testing.T) {
	book := OrderBook{
		BuyOrders: []common.Order{MBOrder4, MBOrder5, MBOrder6, MBOrder7},
	}

	fakeTrade := common.Trade{
		TakerOrdID: 1,
		MakerOrdID: 1,
		Price:      70,
		Quantity:   1,
	}
	trades := book.processLimitSell(LSOrder6)

	if len(book.BuyOrders) != 3 {
		t.Fatal("order not process")
	} else if book.BuyOrders[0] != MBOrder4 {
		t.Fatal("order 0 is not same")
	} else if book.BuyOrders[1] != MBOrder5 {
		t.Fatal("order 1 is not same")
	} else if book.BuyOrders[2] != MBOrder6 {
		t.Fatal("order 2 is not same")
	} else if len(trades) != 1 {
		t.Fatal("invalid trade length")
	} else if trades[0] != fakeTrade {
		t.Fatal("invalid trade")
	}
}

func TestProcessLimitSellAndBuyOrderRemaining(t *testing.T) {
	book := OrderBook{
		BuyOrders: []common.Order{MBOrder4, MBOrder5, MBOrder6, {
			Name:     "test7",
			ReqOrdID: 1,
			Price:    70,
			Quantity: 2,
			Type:     0,
			Side:     0,
		}},
	}

	fakeTrade := common.Trade{
		TakerOrdID: 1,
		MakerOrdID: 1,
		Price:      70,
		Quantity:   1,
	}
	trades := book.processLimitSell(LSOrder6)

	if len(book.BuyOrders) != 4 {
		t.Fatal("order not process")
	} else if book.BuyOrders[0] != MBOrder4 {
		t.Fatal("order 0 is not same")
	} else if book.BuyOrders[1] != MBOrder5 {
		t.Fatal("order 1 is not same")
	} else if book.BuyOrders[2] != MBOrder6 {
		t.Fatal("order 2 is not same")
	} else if book.BuyOrders[3] != MBOrder7 {
		t.Fatal("order 3 is not same")
	} else if len(trades) != 1 {
		t.Fatal("invalid trade length")
	} else if trades[0] != fakeTrade {
		t.Fatal("invalid trade")
	}
}

func TestProcessLimitSellAndSellOrderRemaining(t *testing.T) {
	book := OrderBook{
		BuyOrders: []common.Order{MBOrder4, MBOrder5, MBOrder6, MBOrder7},
	}

	trade1, trade2, trade3, trade4 := common.Trade{
		TakerOrdID: 1,
		MakerOrdID: 1,
		Price:      70,
		Quantity:   1,
	}, common.Trade{
		TakerOrdID: 1,
		MakerOrdID: 1,
		Price:      60,
		Quantity:   1,
	}, common.Trade{
		TakerOrdID: 1,
		MakerOrdID: 1,
		Price:      50,
		Quantity:   1,
	}, common.Trade{
		TakerOrdID: 1,
		MakerOrdID: 1,
		Price:      40,
		Quantity:   1,
	}

	remainingOrder := common.Order{
		Name:     "test1",
		ReqOrdID: 1,
		Price:    40,
		Quantity: 1,
		Type:     1,
		Side:     1,
	}

	trades := book.processLimitSell(common.Order{
		Name:     "test1",
		ReqOrdID: 1,
		Price:    40,
		Quantity: 5,
		Type:     1,
		Side:     1,
	})

	fmt.Println(book.SellOrders[0])

	if len(book.BuyOrders) != 0 {
		t.Fatal("order not process")
	} else if len(trades) != 4 {
		t.Fatal("invalid trade length")
	} else if trades[0] != trade1 {
		t.Fatal("trade 1 not same")
	} else if trades[1] != trade2 {
		t.Fatal("trade 2 not same")
	} else if trades[2] != trade3 {
		t.Fatal("trade 3 not same")
	} else if trades[3] != trade4 {
		t.Fatal("trade 4 not same")
	} else if len(book.SellOrders) != 1 {
		t.Fatal("invalid buy order length")
	} else if book.SellOrders[0] != remainingOrder {
		t.Fatal("sell order 1 not same")
	}
}

func TestProcessLimitSellToLimit(t *testing.T) {
	book := OrderBook{
		BuyOrders: []common.Order{MBOrder4, MBOrder5, MBOrder6, MBOrder7},
	}

	trade1, trade2 := common.Trade{
		TakerOrdID: 1,
		MakerOrdID: 1,
		Price:      70,
		Quantity:   1,
	}, common.Trade{
		TakerOrdID: 1,
		MakerOrdID: 1,
		Price:      60,
		Quantity:   1,
	}

	remainingOrder := common.Order{
		Name:     "test1",
		ReqOrdID: 1,
		Price:    60,
		Quantity: 2,
		Type:     1,
		Side:     1,
	}

	trades := book.processLimitSell(common.Order{
		Name:     "test1",
		ReqOrdID: 1,
		Price:    60,
		Quantity: 4,
		Type:     1,
		Side:     1,
	})

	fmt.Println(book.BuyOrders)
	if len(book.BuyOrders) != 2 {
		t.Fatal("order not process")
	} else if len(trades) != 2 {
		t.Fatal("invalid trade length")
	} else if trades[0] != trade1 {
		t.Fatal("trade 1 not same")
	} else if trades[1] != trade2 {
		t.Fatal("trade 2 not same")
	} else if len(book.SellOrders) != 1 {
		t.Fatal("invalid buy order length")
	} else if book.SellOrders[0] != remainingOrder {
		t.Fatal("sell order 1 not same")
	}
}
