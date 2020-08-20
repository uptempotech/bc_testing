package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/uptempotech/bc_testing/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var client proto.BlockchainClient

func main() {
	addFlag := flag.Bool("add", false, "add new block")
	listFlag := flag.Bool("list", false, "list the blockchain")
	flag.Parse()

	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	client = proto.NewBlockchainClient(conn)

	if *addFlag {
		addBlock()
	}

	if *listFlag {
		getBlockchain()
	}
}

func addBlock() {
	block, err := client.AddBlock(context.Background(), &proto.AddBlockRequest{
		Data: time.Now().String(),
	})
	if err != nil {
		log.Fatalf("Unable to add block to blockchain: %v", err)
	}
	log.Printf("New Block Added with hash: %s\n", block.Hash)
}

func getBlockchain() {
	blocks, err := client.GetBlockchain(context.Background(), &proto.GetBlockchainRequest{})
	if err != nil {
		log.Fatalf("Error fetching blockchain: %v", err)
	}

	counter := 0
	fmt.Println("Blockchain blocks")
	for _, b := range blocks.Blocks {
		fmt.Printf("Block #%d: %s\n", counter, b.Hash)
		fmt.Printf("   Timestamp: %d\n", b.Timestamp)
		fmt.Printf("   Height: %d\n", b.Height)
		fmt.Printf("   Data: %s\n", b.Data)
		fmt.Printf("   Previous Hash: %s\n", b.PrevBlockHash)
		fmt.Printf("   Nonce: %d\n", b.Nonce)

		counter++
	}
}
