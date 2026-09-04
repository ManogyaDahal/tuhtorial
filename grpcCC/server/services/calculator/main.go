package main


func main() { 
	grpcServer := NewGrpcServer(":5050")
	grpcServer.Run()
}
