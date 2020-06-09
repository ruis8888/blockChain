package block

import "time"

// 区块数据
type BlockData struct {
	Version        int
	HashPrevBlock  Hash
	HashMerkleRoot Hash
	Time           time.Time
	Bits           int
	Nonce          int
	Txs       string
	TxCounter int
	HashCurr  Hash
}

//序列化
func BlockSerialize(b Block) ([]byte, error) {

	return nil, nil
}


