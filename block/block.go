package block

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Hash = string
const nodeVersion = 0

// 区块主体
type Block struct {
	Header    BlockHeader
	Txs       string // 交易列表
	txCounter int    // 交易计数器
	HashCurr  Hash   // 当前区块哈市值缓存
}

// 区块头
type BlockHeader struct {
	version        int
	HashPrevBlock  Hash      // 前一个区块的 Hash
	hashMerkleRoot Hash      // 默克尔树的哈希节点
	Time           time.Time // 区块的创建时间
	bits           int       // 难度相关
	nonce          int       // 挖矿相关
}

// 设置当前区块 hash
func (b *Block) SetHashCurr() *Block {
	// 生成头信息的拼接字符串
	headerStr := b.Header.Stringify()
	// 计算 hash 值
	b.HashCurr = fmt.Sprintf("%x", sha256.Sum256([]byte(headerStr)))

	return b
}

// 头信息的字符串化
func (bh *BlockHeader) Stringify() string {
	return fmt.Sprintf("%d%s%s%d%d%d",
		bh.version,
		bh.HashPrevBlock,
		bh.hashMerkleRoot,
		bh.Time.UnixNano(), // 得到时间戳，nano 级别
		bh.bits,
		bh.nonce,
	)
}

// 构造区块
func NewBlock(prevHash Hash, txs string) *Block {
	// 实例化Block
	b := &Block{
		Header:    BlockHeader{
			version:       nodeVersion,
			HashPrevBlock: prevHash, // 设置前面的区块哈希
			Time:          time.Now(),
		},
		Txs:       txs, // 设置数据
		txCounter: 1,   // 计数交易
	}
	// 计算设置当前区块的哈希
	b.SetHashCurr()
	return b
}