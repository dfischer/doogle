package main

import (
	"log"
	"net"
	"context"

	pb "github.com/mathetake/doogle/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"github.com/mathetake/doogle/node"
	"time"
)

const (
	port = ":50051"
	address = "localhost:50051"
)

func main() {
	go func () {runServer()}()
	runClient()
}


// for testing purpose
func runClient() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewDooglleClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Ping(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}

func runServer() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDooglleServer(s, &node.Node{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
