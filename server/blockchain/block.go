package blockchain

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"encoding/hex"
	"log"
	"strconv"
	"time"
)

// Block defines the structure of a block.
type Block struct {
	Timestamp     int64
	Height        int64
	PrevBlockHash string
	Data          string
	Nonce         int64
	Hash          string
}

func (b *Block) serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}

	return res.Bytes()
}

func intToBytes(num int64) []byte {
	s := strconv.FormatInt(num, 10)

	return []byte(s)
}

func bytesToInt(data []byte) int64 {
	var i int64
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &i)

	return i
}

func newBlock(data, prevBlockHash string, height int64) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		Height:        height,
		Data:          data,
		PrevBlockHash: prevBlockHash,
		Nonce:         0,
		Hash:          "",
	}

	pow := newProof(block)
	nonce, hash := pow.Run()

	block.Hash = hex.EncodeToString(hash[:])
	block.Nonce = nonce

	return block
}
