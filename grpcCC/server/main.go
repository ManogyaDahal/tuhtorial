package main

import (
	"context"
	"log"
	"net"

	pb "ManogyaDahal/server/proto/gen/Add"

	"google.golang.org/grpc"
)

type AddNumberServer struct { 
	pb.UnimplementedAddNumberServiceServer
}

func (s *AddNumberServer) AddNumber(ctx context.Context, req *pb.AddNumberRequest )(*pb.AddNumberResponse, error){ 
	var sum int64 = 0 
	for _, value :=	range req.InputNum { 
		sum = sum + value
	}

	log.Println("sum:", sum)

	return &pb.AddNumberResponse{
		Sum: sum,	
	}, nil
}

func main() {
	// specifying the port from which we want to listen to client's request
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("error occured while listening, err:", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterAddNumberServiceServer(grpcServer, &AddNumberServer{})
	
	log.Println("Server is running on port :8080")
	grpcServer.Serve(lis)
}
