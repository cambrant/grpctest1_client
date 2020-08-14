package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "github.com/cambrant/grpctest1_pb/go"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("unable to connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMessageClient(conn)

	text := "blah blah"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SendMessage(ctx, &pb.Msg{Text: text})
	if err != nil {
		log.Fatalf("could not send message: %v", err)
	}
	log.Printf("Message: %s", r.GetText())
}
