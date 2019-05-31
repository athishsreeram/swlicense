package grpc

import (
	"context"
	"net"
	"swlicense/output/eventstore-read-output/proto"
	v1 "swlicense/output/eventstore-read-output/service/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func RunServer(ctx context.Context, port string) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	v1API := v1.NewEventStoreServiceServer()

	proto.RegisterEventStoreServiceServer(srv, v1API)
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}

}
