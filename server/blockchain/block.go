package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/gob"
	"encoding/hex"
	"log"
	"strconv"
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

func (b *Block) setHash() {
	data := bytes.Join(
		[][]byte{
			[]byte(b.Data),
			[]byte(b.PrevBlockHash),
			intToBytes(b.Timestamp),
			intToBytes(b.Height),
			intToBytes(b.Nonce),
		},
		[]byte{},
	)
	hash := sha256.Sum256([]byte(data))
	b.Hash = hex.EncodeToString(hash[:])
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
