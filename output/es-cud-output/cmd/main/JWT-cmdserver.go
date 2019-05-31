package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	cfg "swlicense/output/es-cud-output/config"
	grpc "swlicense/output/es-cud-output/server/grpc"
	rest "swlicense/output/es-cud-output/server/rest"
)

// RunServer runs gRPC server and HTTP gateway
func main() {
	ctx := context.Background()

	var configPath string

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	flag.StringVar(&configPath, "config-path", dir+"/config-dev.json", "Config Path")

	flag.Parse()

	if len(configPath) == 0 {
		fmt.Errorf("invalid Config Path: '%s'", configPath)
	}

	cfg.Conf = cfg.GetConfig(configPath)

	log.Println(cfg.Conf)

	go func() {
		_ = rest.RunServer(ctx, cfg.Conf.GRPCPort, cfg.Conf.HTTPPort)
	}()

	grpc.RunServer(ctx, cfg.Conf.GRPCPort)

}
