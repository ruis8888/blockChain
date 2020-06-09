package blockchain

import (
	"blockChain/block"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"

	"time"
)

type BlockChain struct {
	lastHash block.Hash // 最后一个区块的哈希
	db *leveldb.DB
}
//区块链的构造方法
func NewBlockchain(db *leveldb.DB) *BlockChain {
		bc := &BlockChain{
		db:   db,
	}
	return bc
}
// 添加区块
// 提供区块的数据，目前是字符串
func (bc *BlockChain) AddBlock(txs string) *BlockChain {
	//构造区块
	b := block.NewBlock(bc.lastHash,txs)
	//把区块加入到链存储结构中中

	// 将最后的区块哈希设置为当前区块
	bc.lastHash = b.HashCurr


	return bc
}

//添加创世区块
func (bc *BlockChain) AddGensisBlock() *BlockChain {
	//检验是否能添加区块
	if bc.lastHash!="" {
		return bc
	}
	return bc.AddBlock("this is gensis block")
}

//迭代的方式展示区块
func (bc *BlockChain) Iterate() {
	// 最后的哈希
	for hash := bc.lastHash; hash != ""; {
		b := bc.blocks[hash]
		fmt.Println("HashCurr:", b.HashCurr)
		fmt.Println("TXs:", b.Txs)
		fmt.Println("Time:", b.Header.Time.Format(time.UnixDate))
		fmt.Println("HashPrev:", b.Header.HashPrevBlock)
		fmt.Println()

		hash = b.Header.HashPrevBlock
	}
}


