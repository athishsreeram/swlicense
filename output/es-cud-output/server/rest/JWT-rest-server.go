package rest

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	gw "output/es-cud-output/proto"
)

func RunServer(ctx context.Context, grpcport string, httpport string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterJWTServiceHandlerFromEndpoint(ctx, mux, "localhost:"+grpcport, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":"+httpport, mux)
}
