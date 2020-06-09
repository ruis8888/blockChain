package main

import (
	"blockChain/block"
	"bytes"
	"encoding/gob"
	"fmt"
)

type Block struct {
	Txs string
	HashCurr string
}

func main() {
	b := Block{
		Txs:      "first block",
		HashCurr: "1111222233334444",
	}

	var bb bytes.Buffer
	encoder := gob.NewEncoder(&bb)
	encoder.Encode(b)
	fmt.Println(bb.Bytes(),bb.String())
	result := bb.Bytes()

	//解码
	var bbb bytes.Buffer
	//先写进缓冲中
	bbb.Write(result)
	decoder := gob.NewDecoder(&bbb)
	b1 := block.Block{}
	decoder.Decode(&b1)
	fmt.Println(b1)
}
