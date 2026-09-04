package main

import (
	genpb "client/proto/gen"
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


const domain = "localhost:5050"
func main(){
	conn, err := grpc.NewClient(domain, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil { 
		log.Fatal("Error occured while creating client", err)
	}
	defer conn.Close()

	ctx := context.Background()

	client := genpb.NewFibonacciClient(conn)
	req := &genpb.FibonacciRequest{ N: 10 }
	stream, err := client.GenerateFibonacci(ctx, req)
	if err != nil { 
		log.Fatal("error occured", err)
	}

	for { 
		resp, err := stream.Recv()
		if err != nil { 
			if err == io.EOF { 
				log.Println("End of file")
				break
			}
			log.Fatal(err)
		}
		fmt.Printf("%d ,",resp.Num)
	}


	//-----------------------------------------------------

	strm , err := client.SendNumbers(ctx)
	if err != nil { 
		log.Fatal(err)
	}

	for num := range 9 { 
		err := strm.Send(&genpb.NumberRequest{Number: int32(num)})
		if err != nil { 
			log.Fatal("An error occured", err)
		}
		time.Sleep(time.Second)
	}

	resp , err:= strm.CloseAndRecv()
	if err != nil { 
		log.Println("error occured colse and recv", err)
	}
	log.Println("server response after stream", resp.Num)
}
