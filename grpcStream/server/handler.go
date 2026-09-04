package main

import (
	pb "server/proto/gen"
)

// Handler struct
type FibonacciHandler struct {
	pb.UnimplementedFibonacciServer
	svc fiboSvc	
}

// Initializes new Handler.
func NewHandler(svc fiboSvc) *FibonacciHandler { 
	return &FibonacciHandler{ 
		svc:  svc,
	}
}

// Generating fibonacci handler
func (h *FibonacciHandler) GenerateFibonacci (req *pb.FibonacciRequest, stream pb.Fibonacci_GenerateFibonacciServer) error { 
	return h.svc.GrenerateFibonacciSvc(req, stream)
}

// returning sum
func (h *FibonacciHandler) SendNumbers (stream pb.Fibonacci_SendNumbersServer)error{ 
	return h.svc.SendNumbersSvc(stream)
}
