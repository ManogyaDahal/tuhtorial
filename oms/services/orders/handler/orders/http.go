package handler

import (
	"net/http"

	"ManogyaDahal/oms/services/common/genproto/orders"
	"ManogyaDahal/oms/services/orders/types"
	"ManogyaDahal/oms/util"
)

type OrdersHttpHandler struct {
	ordersService types.OrderService
}

func NewHttpOrdersHandler(orderService types.OrderService) *OrdersHttpHandler {
	handler := &OrdersHttpHandler{
		ordersService: orderService,
	}

	return handler
}

func (h *OrdersHttpHandler) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("POST /orders", h.CreateOrder)
}

func (h *OrdersHttpHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req orders.CreateOrderRequest
	err := util.ParseJson(r, &req)
	if err != nil {
		// util.WriteError(w, http.StatusBadRequest, err)
		return
	}

	order := &orders.Order{
		OrderId:    10,
		CustomerId: req.GetCustomerId(),
		ProductId:  req.GetProductId(),
		Quantity:   req.Quantity,
	}
	err = h.ordersService.CreateOrder(r.Context(), order)
	if err != nil { 
		util.WriteError(w, http.StatusInternalServerError, err)
		return 
	}

	res := &orders.CreateOrderResponse{Status: "success"}
	util.WriteJson(w, http.StatusOK, res)	
}
