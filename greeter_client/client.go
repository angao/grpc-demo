package main

import (
	"context"
	"log"
	"os"

	pb "github.com/angao/grpc-demo/helloworld"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	gc := pb.NewGreeterClient(conn)
	sc := pb.NewSenderClient(conn)

	name := "World"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	reply, err := gc.SayHello(context.Background(), &pb.HelloRequest{
		Name: name,
	})
	if err != nil {
		log.Fatalf("failed to greet: %v", err)
	}
	log.Printf("Greeting: %s", reply.Message)

	rp, err := sc.Send(context.Background(), &pb.HelloRequest{
		Name: "send1",
	})

	if err != nil {
		log.Fatalf("failed to greet: %v", err)
	}
	log.Printf("Greeting: %s", rp.Message)
}
