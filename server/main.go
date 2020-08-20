package main

import (
	"context"
	"log"
	"net"

	"github.com/uptempotech/bc_testing/proto"
	"github.com/uptempotech/bc_testing/server/blockchain"
	"google.golang.org/grpc"
)

type blockServer struct {
	Blockchain *blockchain.Blockchain
}

func (server *blockServer) AddBlock(_ context.Context, in *proto.AddBlockRequest) (*proto.AddBlockResponse, error) {
	block := server.Blockchain.AddBlock(in.GetData())
	return &proto.AddBlockResponse{
		Hash: block.Hash,
	}, nil
}

func (server *blockServer) GetBlockchain(_ context.Context, in *proto.GetBlockchainRequest) (*proto.GetBlockchainResponse, error) {
	resp := new(proto.GetBlockchainResponse)
	for _, b := range server.Blockchain.Blocks {
		resp.Blocks = append(resp.Blocks, &proto.Block{
			Timestamp:     b.Timestamp,
			Height:        b.Height,
			Data:          b.Data,
			PrevBlockHash: b.PrevBlockHash,
			Nonce:         b.Nonce,
			Hash:          b.Hash,
		})
	}
	return resp, nil
}

func main() {
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("Error starting listener: %v", err)
	}

	srv := grpc.NewServer()
	proto.RegisterBlockchainServer(srv, &blockServer{
		Blockchain: blockchain.NewBlockChain(),
	})
	srv.Serve(listener)
}
