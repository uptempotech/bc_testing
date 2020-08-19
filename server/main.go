package main

import (
	"context"
	"log"
	"net"

	"github.com/uptempotech/bc_testing/proto"
	"google.golang.org/grpc"
)

type blockServer struct{}

func (server *blockServer) AddBlock(_ context.Context, in *proto.AddBlockRequest) (*proto.AddBlockResponse, error) {
	return new(proto.AddBlockResponse), nil
}

func (server *blockServer) GetBlockchain(_ context.Context, in *proto.GetBlockchainRequest) (*proto.GetBlockchainResponse, error) {
	return new(proto.GetBlockchainResponse), nil
}

func main() {
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("Error starting listener: %v", err)
	}

	srv := grpc.NewServer()
	proto.RegisterBlockchainServer(srv, &blockServer{})
	srv.Serve(listener)
}
