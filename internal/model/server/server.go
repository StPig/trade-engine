package server

import (
	"context"
	"fmt"
	"trade-engine/internal/model/common"
	"trade-engine/internal/model/engine"
	"trade-engine/pkg/protocol"
)

type Server struct{}

// Create order book
var book = engine.OrderBook{
	BuyOrders:  make([]common.Order, 0, 100),
	SellOrders: make([]common.Order, 0, 100),
}

// Create new order to process order via gRPC
func (s *Server) CreateOrder(ctx context.Context, req *protocol.CreateOrderRequest) (*protocol.ExecutionReport, error) {
	fmt.Println("create order")

	order := common.Order{
		Name:     req.GetName(),
		ReqOrdID: req.GetReqOrdID(),
		Price:    req.GetPrice(),
		Quantity: req.GetQuantity(),
		Type:     int(req.GetOrderType()),
		Side:     int(req.GetOrderSide()),
	}

	trades := book.Process(order)

	fmt.Println(trades)
	status := "pending"
	if len(trades) > 0 {
		status = "success"
	}

	res := &protocol.ExecutionReport{
		Name:     order.Name,
		ReqOrdID: order.ReqOrdID,
		Status:   status,
	}
	return res, nil
}
