package v1

import (
	"output/es-cud-output/client/nats/pub"
	"output/es-cud-output/proto"

	cfg "output/es-cud-output/config"

	guuid "github.com/google/uuid"

	"context"
	"log"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

// jwtServer is implementation of proto.JWTServer proto interface
type jwtServer struct {
}

// NewJWTServiceServer creates {$apiname}} service
func NewJWTServiceServer() proto.JWTServiceServer {
	return &jwtServer{}
}

// Create new todo task
func (s *jwtServer) Create(ctx context.Context, req *proto.JWTRequest) (*proto.JWTResponse, error) {
	uuid := guuid.New().String()
	pub.Send(cfg.Conf.NATSSubj, "JWTCreateCommand", "command", req, uuid)

	var resp *proto.JWTResponse

	resp = &proto.JWTResponse{
		V1: 1,

		JWTToken: nil,
	}

	log.Println("called out")
	log.Println(resp)
	for resp == nil {
	}
	return resp, nil
}

// Create new todo task
func (s *jwtServer) Update(ctx context.Context, req *proto.JWTRenewRequest) (*proto.JWTRenewResponse, error) {
	uuid := guuid.New().String()
	pub.Send(cfg.Conf.NATSSubj, "JWTUpdateCommand", "command", req, uuid)

	var resp *proto.JWTRenewResponse

	resp = &proto.JWTRenewResponse{
		V1: 1,

		JWTToken: nil,
	}

	log.Println("called out")
	log.Println(resp)
	for resp == nil {
	}
	return resp, nil
}
