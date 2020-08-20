package blockchain

import "time"

func newBlock(data, prevBlockHash string, height int64) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		Height:        height,
		Data:          data,
		PrevBlockHash: prevBlockHash,
		Nonce:         0,
		Hash:          "",
	}

	return block
}
