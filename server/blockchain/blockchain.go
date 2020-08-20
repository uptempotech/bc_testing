package blockchain

import "fmt"

// Blockchain defines the structure of the blockchain itself.
type Blockchain struct {
	Blocks []*Block
}

// AddBlock will a block to the blockchain.
func (bc *Blockchain) AddBlock(data string) *Block {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]

	block := newBlock(data, prevBlock.Hash, prevBlock.Height+1)

	bc.Blocks = append(bc.Blocks, block)

	return block
}

// NewBlockChain creates new blockchain.
func NewBlockChain() *Blockchain {
	fmt.Println("Initializing new blockchain with Genesis block:")
	return &Blockchain{[]*Block{newGenesisBlock()}}
}

func newGenesisBlock() *Block {
	return newBlock("Genesis Block", "", 0)
}
