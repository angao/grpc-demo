package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc/reflection"

	pb "github.com/angao/grpc-demo/helloworld"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("receive message: %s", in.Name)
	return &pb.HelloReply{
		Message: "Hello " + in.Name,
	}, nil
}

type sender struct{}

func (s *sender) Send(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message: "Send " + in.Name,
	}, nil
}

func main() {
	ls, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterGreeterServer(s, &server{})
	pb.RegisterSenderServer(s, &sender{})
	reflection.Register(s)

	if err := s.Serve(ls); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
