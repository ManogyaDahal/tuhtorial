package main

import (
	"context"
	"log"
	"time"

	pb "ManogyaDahal/client/proto/gen/Add"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func main() {
	addr := "localhost:8080"
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Failed connection, err", err)
	}
	defer conn.Close()

	// defining client 
	client := pb.NewAddNumberServiceClient(conn)
	ctx, cancel:= context.WithTimeout(context.Background(), time.Second * 10)
	defer cancel()
	
	req := &pb.AddNumberRequest{ 
		InputNum: []int64{2, 2, 2},
	}
	resp, err := client.AddNumber(ctx, req) //pulling Rpc method
	if err != nil { 
		log.Println("some error occured", err)
	}
	log.Println("obtained Sum is:", resp.Sum)
}
