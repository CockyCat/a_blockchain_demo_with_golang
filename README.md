# Step by one step to develop a blockchain project

## Define struct of Block
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