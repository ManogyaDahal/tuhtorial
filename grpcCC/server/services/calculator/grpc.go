package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

type GrpcServer struct {
	addr string
}

// Used to initialize a new Grpc Server.
// Includes Run function to runn the server
// from the main funciton
func NewGrpcServer(addr string) *GrpcServer {
	return &GrpcServer{
		addr: addr,
	}
}

// Responsible for running the server
func (s *GrpcServer) Run() error {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatal("Error occured while listening", err)
	}
	grpcServer := grpc.NewServer()
	log.Println("GrpcServer is running on ",s.addr)
	// defining the service

	return grpcServer.Serve(listener)
}
