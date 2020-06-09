package main

import (
	"blockChain/block"
	"fmt"
)

func main() {
	b := block.NewBlock("", "Gensis block")
	fmt.Println(b)
}
