package main

import (
	"log"
	"net"

	pb "github.com/HyanSource/grpctest/api"
	"github.com/HyanSource/grpctest/handler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Println(err)
		return
	}

	s := grpc.NewServer()
	pb.RegisterAPIServiceServer(s, &handler.APIServer{})
	reflection.Register(s)
	log.Println("Server Start")

	if err := s.Serve(l); err != nil {
		log.Println(err)
		return
	}
	log.Println("Server Over")
}
