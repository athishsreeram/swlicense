package v1

import (
	"context"
	"log"
	"output/es-read-output/custom"
	"output/es-read-output/domain"
	"output/es-read-output/proto"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

// jwtServer is implementation of proto.JWTServer proto interface
type jwtServer struct {
}

// NewJWTServiceServer creates JWT service
func NewJWTServiceServer() proto.JWTServiceServer {
	return &jwtServer{}
}

// Create new todo task
func (s *jwtServer) Read(ctx context.Context, req *proto.JWTVerifyRequest) (*proto.JWTVerifyResponse, error) {
	// TO-DO

	var jWTTokenResp *proto.JWTToken

	jWTTokenResp = domain.ConvertJWTTokenDomains2JWTToken(domain.ReadJWTTokenDomains(req.JwtToken))

	resp := &proto.JWTVerifyResponse{
		V1: 1,

		Status: custom.VerifyJWT(jWTTokenResp.Token),
	}
	log.Println(jWTTokenResp)
	log.Println(resp)
	return resp, nil
}
