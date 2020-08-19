package blockchain

func newBlock(data, prevBlockHash string) *Block {
	block := &Block{
		Data:          data,
		PrevBlockHash: prevBlockHash,
	}

	return block
}
