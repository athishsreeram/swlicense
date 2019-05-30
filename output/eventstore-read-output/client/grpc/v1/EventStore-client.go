package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"

	v1 "output/eventstore-read-output/proto"
)

func main() {
	// Get configuration
	address := flag.String("server", "localhost:4040", "gRPC server in format host:port")
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := v1.NewEventStoreServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Call ReadAll

	req0 := v1.ReadAllRequest{

		V1: 1,
	}
	res0, err := c.ReadAll(ctx, &req0)
	if err != nil {
		log.Fatalf("Call ReadAll failed: %v", err)
	}
	log.Printf("Call ReadAll result: &lt;%+v>\n\n", res0)

	// Call Read

	req1 := v1.ReadRequest{

		V1:   1,
		Uuid: 1,
	}
	res1, err := c.Read(ctx, &req1)
	if err != nil {
		log.Fatalf("Call Read failed: %v", err)
	}
	log.Printf("Call Read result: &lt;%+v>\n\n", res1)

}
