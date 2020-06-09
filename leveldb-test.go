package main

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
)

func main() {
	//打开
	dbPath := "testdb"
	db, err := leveldb.OpenFile(dbPath, nil)
	if err!=nil {
		log.Fatal(err)
	}
	key := "ruis 8888"
	// 设置
	if err := db.Put([]byte(key), []byte("blockChain Demo"), nil); err!=nil{
		log.Fatal(err)
	}
	log.Println("put success!")

	// 读取key
	data, err := db.Get([]byte(key), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data, string(data))
}
