# Step by one step to develop a blockchain project

## 1. Define struct of Block
```
type Block struct {
	Height       int64
	PreBlockHash []byte
	Transactions []byte
	Timestamp    int64
	Hash         []byte
	Nonce        int64
}
```
## 2. Declare a func that named GenesisBlock to create genesis block
```
func GenesisBlock(data string) *Block {
	return NewBlock(
		data,
		1,
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	)
}
```
## 3. Define struct of Blockchain that includes a group of blocks
```
type Blockchain struct {
	Blocks []*Block
}
```
## 4. To create genesis blockchain 
```
func GenesisBlockChain() *Blockchain {
	gensisBlock := GenesisBlock("Geneis Block")
	return &Blockchain{[]*Block{gensisBlock}}
}
```