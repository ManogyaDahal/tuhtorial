package main

import (
	"log"
	"net"
	pb "server/proto/gen"

	"google.golang.org/grpc"
)

type server struct { 
	addr string
	handler *FibonacciHandler 
}

func NewServer(addr string, handler *FibonacciHandler) *server { 
	return &server { 
		addr: addr, 
		handler: handler,
	}
}

func (s *server)Run() error { 
	lis , err := net.Listen("tcp", s.addr)
	if err != nil { 
		return err
	}
	grpcServer := grpc.NewServer()
	// Registering services
	// passed handler because it contains the UnimplementedFibonacciServer
	pb.RegisterFibonacciServer(grpcServer, s.handler) 
	log.Println("Server running on port:",s.addr)

	return grpcServer.Serve(lis)
}
