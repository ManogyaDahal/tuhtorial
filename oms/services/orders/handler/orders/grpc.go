package handler

import (
	"context"

	"ManogyaDahal/oms/services/common/genproto/orders"
	"ManogyaDahal/oms/services/orders/types"

	"google.golang.org/grpc"
)

type OrdersGrpcHandler struct {
	ordersService types.OrderService
	orders.UnimplementedOrderServiceServer
}

func NewGrpcOrderService(grpc *grpc.Server, ordersService types.OrderService) {
	grpcHandler := &OrdersGrpcHandler{
		ordersService:  ordersService,
	}

	orders.RegisterOrderServiceServer(grpc, grpcHandler)
}


func (h *OrdersGrpcHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	// mock data for order
	order := &orders.Order{
		OrderId:    43,
		CustomerId: 3,
		ProductId:  2,
		Quantity:   4,
	}
	err := h.ordersService.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	res := &orders.CreateOrderResponse{
		Status: "success",
	}

	return res, nil
}
