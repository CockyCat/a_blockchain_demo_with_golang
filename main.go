package main

import (
	"fmt"

	"gary.com/myweb/pubchain/block"
)

func main() {
	block := block.NewBlock("Geneis Block", 10, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	fmt.Println(block)
}
