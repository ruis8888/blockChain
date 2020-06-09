package blockchain

import (
	"blockChain/block"
	
	"github.com/syndtr/goleveldb/leveldb"

)
type BlockChain struct {
	lastHash block.Hash  // 最后一个区块的哈希
	db       *leveldb.DB // leveldb 的连接
}




