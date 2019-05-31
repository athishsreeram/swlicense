package v1

import (
	"context"
	"log"
	"swlicense/output/eventstore-read-output/domain"
	"swlicense/output/eventstore-read-output/proto"
	"strconv"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

// eventstoreServer is implementation of proto.EventStoreServer proto interface
type eventstoreServer struct {
}

// NewEventStoreServiceServer creates EventStore service
func NewEventStoreServiceServer() proto.EventStoreServiceServer {
	return &eventstoreServer{}
}

// Create new todo task
func (s *eventstoreServer) ReadAll(ctx context.Context, req *proto.ReadAllRequest) (*proto.ReadAllResponse, error) {
	// TO-DO

	var list []*proto.EventStore
	for _, eventStoresResp := range domain.ReadAllEventStores() {
		eventStore := domain.ConvertEventStores2EventStore(eventStoresResp)
		list = append(list, eventStore)
	}

	resp := &proto.ReadAllResponse{
		V1: 1,

		EventStores: list,
	}
	log.Println(resp)
	return resp, nil
}

// Create new todo task
func (s *eventstoreServer) Read(ctx context.Context, req *proto.ReadRequest) (*proto.ReadResponse, error) {
	// TO-DO

	var eventStore *proto.EventStore

	eventStore = domain.ConvertEventStores2EventStore(domain.ReadEventStores(req.Uuid))

	resp := &proto.ReadResponse{
		V1: 1,

		EventStore: eventStore,
	}
	log.Println(resp)
	return resp, nil
}
