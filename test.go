package main

import (
	"blockChain/blockchain"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"log"

)

func main() {
	//b := block.NewBlock("", "Gensis block")
	//fmt.Println(b)

	//区块链的测试
	//bc := blockchain.NewBlockChain()
	//bc.AddGensisBlock()
	//bc.AddBlock("First block").AddBlock("second block")
	//bc.Iterate()

	//数据库连接
	dbPath := "data"
	db, err := leveldb.OpenFile(dbPath, nil)
	if err!=nil {
		log.Fatal(err)
	}
	//释放数据库连接
	defer db.Close()
	//区块链测试
	bc := blockchain.NewBlockchain(db)
	fmt.Println(bc)
}
