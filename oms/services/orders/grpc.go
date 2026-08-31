package main

import (
	"log"
	"net"

	handler "ManogyaDahal/oms/services/orders/handler/orders"
	"ManogyaDahal/oms/services/orders/service"

	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{
		addr: addr,
	}
}

func (g *gRPCServer) Run() error {
	listener, err := net.Listen("tcp", g.addr)
	if err != nil {
		log.Fatal("Error occured while listening. err:", err)
	}
	grpcServer := grpc.NewServer()
	// need to register grpc services
	orderService := service.NewOrderService()
	handler.NewGrpcOrderService(grpcServer, orderService)

	log.Println("grpc Server running on port: ", g.addr)

	return grpcServer.Serve(listener)
}
