package main

import (
	"ManogyaDahal/server/services/calculator/handler"
	"ManogyaDahal/server/services/calculator/service"
)

func main() {
	svc := service.NewService()
	h := handler.NewGrpcHandlers(svc)
	grpcServer := NewGrpcServer(":5050", h)
	grpcServer.Run()
}
