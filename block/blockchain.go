package block

type Blockchain struct {
	Blocks []*Block
}

func (bc *Blockchain) AddBlock2Chain(data string, height int64, preBlockHash []byte) {
	newBlock := NewBlock(data, height, preBlockHash)
	//append a block to chain
	bc.Blocks = append(bc.Blocks, newBlock)
}

func GenesisBlockChain() *Blockchain {
	gensisBlock := GenesisBlock("Geneis Block")

	return &Blockchain{[]*Block{gensisBlock}}
}
