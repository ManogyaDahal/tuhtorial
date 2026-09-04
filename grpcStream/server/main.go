package main



func main(){ 
	svc := NewService()
	handler := NewHandler(svc)

	server := NewServer(":5050", handler)
	server.Run()
}
