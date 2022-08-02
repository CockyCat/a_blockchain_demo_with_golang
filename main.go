package main

import (
	"fmt"

	"gary.com/myweb/a_blockchain_demo_with_golang/block"
)

func main() {
	gensisBlock := block.GenesisBlock("Geneis Block")
	fmt.Println(gensisBlock)
}
