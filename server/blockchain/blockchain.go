package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
)

// Block defines the structure of a block.
type Block struct {
	Hash          string
	PrevBlockHash string
	Data          string
}

// Blockchain defines the structure of the blockchain itself.
type Blockchain struct {
	Blocks []*Block
}

func (b *Block) setHash() {
	hash := sha256.Sum256([]byte(b.PrevBlockHash + b.Data))
	b.Hash = hex.EncodeToString(hash[:])
}
