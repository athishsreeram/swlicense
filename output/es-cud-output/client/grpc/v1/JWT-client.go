package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"

	v1 "output/es-cud-output/proto"
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

	c := v1.NewJWTServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Call Create

	req0 := v1.JWTRequest{

		V1:   1,
		User: user,
	}
	res0, err := c.Create(ctx, &req0)
	if err != nil {
		log.Fatalf("Call Create failed: %v", err)
	}
	log.Printf("Call Create result: &lt;%+v>\n\n", res0)

	// Call Update

	req1 := v1.JWTRenewRequest{

		V1:       1,
		JwtToken: jwtToken,
	}
	res1, err := c.Update(ctx, &req1)
	if err != nil {
		log.Fatalf("Call Update failed: %v", err)
	}
	log.Printf("Call Update result: &lt;%+v>\n\n", res1)

	// Call Read

	req2 := v1.JWTVerifyRequest{

		V1:       1,
		JwtToken: jwtToken,
	}
	res2, err := c.Read(ctx, &req2)
	if err != nil {
		log.Fatalf("Call Read failed: %v", err)
	}
	log.Printf("Call Read result: &lt;%+v>\n\n", res2)

}
