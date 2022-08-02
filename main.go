package main

import (
	"fmt"

	"gary.com/myweb/a_blockchain_demo_with_golang/block"
)

func main() {
	gensisBlockChain := block.GenesisBlockChain()
	fmt.Println(gensisBlockChain)

	gensisBlockChain.AddBlock2Chain(
		"A transaction happened",
		gensisBlockChain.Blocks[len(gensisBlockChain.Blocks)-1].Height+1, //last height add 1
		gensisBlockChain.Blocks[len(gensisBlockChain.Blocks)-1].Hash,
	)

	gensisBlockChain.AddBlock2Chain(
		"A transaction happened",
		gensisBlockChain.Blocks[len(gensisBlockChain.Blocks)-1].Height+1,
		gensisBlockChain.Blocks[len(gensisBlockChain.Blocks)-1].Hash,
	)

	gensisBlockChain.AddBlock2Chain(
		"A transaction happened",
		gensisBlockChain.Blocks[len(gensisBlockChain.Blocks)-1].Height+1,
		gensisBlockChain.Blocks[len(gensisBlockChain.Blocks)-1].Hash,
	)

	gensisBlockChain.AddBlock2Chain(
		"A transaction happened",
		gensisBlockChain.Blocks[len(gensisBlockChain.Blocks)-1].Height+1,
		gensisBlockChain.Blocks[len(gensisBlockChain.Blocks)-1].Hash,
	)
}
