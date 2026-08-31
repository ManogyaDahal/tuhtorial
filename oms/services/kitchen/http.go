package main

import (
	"log"
	"net/http"

	"ManogyaDahal/oms/services/common/genproto/orders"
	handler "ManogyaDahal/oms/services/kitchen/handler/kitchen"
	"ManogyaDahal/oms/util"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{
		addr: addr,
	}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()

	conn := NewGRPCClient(":9000")
	defer conn.Close()

	orderClient := orders.NewOrderServiceClient(conn)
	kitchenHandler := handler.NewHttpKitchensHandler(orderClient)
	kitchenHandler.RegisterRouter(router)

	log.Println("Starting server on ", s.addr)

	return http.ListenAndServe(s.addr, util.EnableCors(router))
}
