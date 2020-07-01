package main

import (
	"github.com/boltdb/bolt"
	"log"
)

//4.引入区块链
type BlockChain struct {
	// 定义一个区块链数组
	//blocks []*Block

	db *bolt.DB
	tail []byte // 存储最后一个区块的哈希

}

const blockChainDb = "blockChain.db"
const blockBucket = "blockBucket"

// 5.定义一个区块链
func NewBlockChain() *BlockChain{

	//return &BlockChain{
	//	blocks: []*Block{genesisBlock},
	//}

	// 最后一个区块的哈希，从数据库中读取的
	var lastHash []byte
	// 1.打开数据库
	db, err := bolt.Open(blockChainDb, 0600, nil)
	//defer db.Close()
	if err != nil {
		log.Panic("打开数据库失败")
	}

	// 将要操作数据库（改写）
	db.Update(func(tx *bolt.Tx) error {
		// 2.找到抽屉bucket,没有就创建
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			// 没有抽屉，我们需要创建
			bucket, err = tx.CreateBucket([]byte(blockBucket))
			if err != nil{
				log.Panic("创建bucket失败")
			}
			// 创建一个创世块，并作为第一个区块添加到区块链中
			genesisBlock := GenesisBlock()
			// 3.写数据
			// hash 作为key，block的字节流作为value，尚未实现
			bucket.Put(genesisBlock.Hash, genesisBlock.Serialize())
			// 最后区块的哈希
			bucket.Put([]byte("lastHashKey"), genesisBlock.Hash)
            // 没有区块则要取创世块的哈希
			lastHash = genesisBlock.Hash
		}else {
			// 有区块则直接取最后区块的哈希
			lastHash = bucket.Get([]byte("lastHashKey"))
		}
		return nil
	})
	return &BlockChain{db, lastHash}
}

// 创世块
func GenesisBlock() *Block{
	return NewBlock("创世块", []byte{})
}

//6.添加区块
func (bc *BlockChain)AddBlock(data string)  {

	db := bc.db
	lastHash := bc.tail // 最后一个区块的哈希

	db.Update(func(tx *bolt.Tx) error {
		// 完成数据添加
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil{
			 log.Panic("bucket 不应该为空，请检查。")
		}

		// a.创建新的区块
		block := NewBlock(data, lastHash)

		// b.添加到区块链db中，保存两部分
		// 区块的 区块哈希 -> 区块序列化字节流
		bucket.Put(block.Hash, block.Serialize())
		// "lastHashKey" -> 最后区块的哈希
		bucket.Put([]byte("lastHashKey"), block.Hash)

		// c.更新内存中的区块链，把最后的小尾巴tail更新一下
		bc.tail = block.Hash
        return nil
	})

	//// 获取最后一个区块
	//lastBlock := bc.blocks[len(bc.blocks)-1]
	//prevHash := lastBlock.Hash
	//
	//// 创建新的区块
	//block := NewBlock(data, prevHash)
	//
	//// 添加到区块链数组中
	//bc.blocks = append(bc.blocks, block)
}