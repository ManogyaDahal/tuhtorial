// contains actual business logic of a service
package main

import (
	"io"
	"log"
	pb "server/proto/gen"
	"time"
) 

type FibonacciService interface {
	GrenerateFibonacciSvc(*pb.FibonacciRequest, pb.Fibonacci_GenerateFibonacciServer) error
	SendNumbersSvc (stream pb.Fibonacci_SendNumbersServer) error
}

type fiboSvc struct{}

func NewService() fiboSvc { 
	return fiboSvc{}
}

func (s *fiboSvc)GrenerateFibonacciSvc(req *pb.FibonacciRequest, stream pb.Fibonacci_GenerateFibonacciServer) error{
	n := req.N
	a, b := 0, 1
	
	for i:=0; i < int(n); i++ { 
		err := stream.Send(&pb.FibonacciResponse{
			Num: int32(a),	
		})
		if err != nil { 
			return nil 
		}
		a,b = b, a+b
		// simulating a delay
		time.Sleep(time.Second)
	}

	return nil
}

func (s *fiboSvc) SendNumbersSvc (stream pb.Fibonacci_SendNumbersServer) error { 
	var sum int32 = 0 
	for { 
		req, err := stream.Recv()
		if err != nil { 
			if err == io.EOF{
				break;
			}
			log.Fatal("error:", err)
		}
		log.Println(req.GetNumber())
		sum = sum + req.Number
	}
	log.Println("FinalSum: ", sum)

	return stream.SendAndClose(&pb.FibonacciResponse{Num: sum})
}
