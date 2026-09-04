package main

import (
	"log"
	"net"

	pb "ManogyaDahal/server/proto/gen/calculator"
	"ManogyaDahal/server/services/calculator/handler"

	"google.golang.org/grpc"
)

type GrpcServer struct {
	addr     string
	handlers *handler.GrpcCalculatorHandler
}

func NewGrpcServer(addr string, handlers *handler.GrpcCalculatorHandler) *GrpcServer {
	return &GrpcServer{
		addr:     addr,
		handlers: handlers,
	}
}

func (s *GrpcServer) Run() error {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatal("Error occured while listening", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterCalculatorServer(grpcServer, s.handlers)
	log.Println("GrpcServer is running on ", s.addr)
	return grpcServer.Serve(listener)
}
