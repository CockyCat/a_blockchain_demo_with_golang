package block

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Height       int64
	PreBlockHash []byte
	Transactions []byte
	Timestamp    int64
	Hash         []byte
}

func (b *Block) CreateHash() {
	heightBytes := IntToBytes(b.Height)
	fmt.Println("Height: ", heightBytes)

	timeString := strconv.FormatInt(b.Timestamp, 10)
	fmt.Println("timeString: ", timeString)
	timeBytes := []byte(timeString)

	blockBytes := bytes.Join([][]byte{heightBytes, b.PreBlockHash, b.Transactions, timeBytes, b.Hash}, []byte{})

	hash := sha256.Sum256(blockBytes)

	b.Hash = hash[:]
}

//New a Block
func NewBlock(transactions string, height int64, preBlockHash []byte) *Block {
	block := &Block{
		Height:       height,
		PreBlockHash: preBlockHash,
		Transactions: []byte(transactions), //字符串强转byte数组
		Timestamp:    time.Now().Unix(),
		Hash:         nil,
	}
	block.CreateHash()
	return block
}

//创世区块
func GenesisBlock(data string) *Block {
	return NewBlock(
		data, 
		1, 
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	)
}
