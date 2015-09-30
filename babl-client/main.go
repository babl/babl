package main

import (
	"log"
	"os"

	pb "github.com/larskluge/babl/protobuf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:4444"
	// address     = "localdocker:8080"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// c := pb.NewGreeterClient(conn)
	c := pb.NewStringUpcaseClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	in := []byte(name)
	r, err := c.IO(context.Background(), &pb.BinRequest{In: in})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Out)
}
