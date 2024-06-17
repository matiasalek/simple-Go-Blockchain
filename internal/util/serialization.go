package util

import (
	"bytes"
	"encoding/gob"
	"log"
	"myBlockchain/internal/blockchain"
)

func Serialize(block *blockchain.Block) []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(block)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}

func DeserializeBlock(d []byte) *blockchain.Block {
	var block blockchain.Block
	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}
	return &block
}
