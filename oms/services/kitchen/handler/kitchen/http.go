package handler

import (
	"net/http"

	"ManogyaDahal/oms/services/common/genproto/orders"
	"ManogyaDahal/oms/util"
)

type KitchensHttpHandler struct {
	orderClient orders.OrderServiceClient
}

func NewHttpKitchensHandler(orderClient orders.OrderServiceClient) *KitchensHttpHandler {
	handler := &KitchensHttpHandler{
		orderClient: orderClient,
	}

	return handler
}

func (h *KitchensHttpHandler) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("POST /orders", h.CreateOrder)
}

func (h *KitchensHttpHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req orders.CreateOrderRequest
	err := util.ParseJson(r, &req)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// kitchen calls orders service over grpc
	res, err := h.orderClient.CreateOrder(r.Context(), &req)
	if err != nil {
		util.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteJson(w, http.StatusOK, res)
}
