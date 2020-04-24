package main

import (
	"crypto/sha256"
	"fmt"
)

//1.定义结构
type Block struct {
	//	前区块哈希
	PrevHash []byte
	//	当前区块哈希
	Hash []byte
	//	数据
	Data []byte

}

//2.创建区块
func NewBlock(data string, prevBloackHash []byte) *Block {
	block := Block{
		PrevHash: prevBloackHash,
		Hash: []byte{}, // 先填空后计算
		Data: []byte(data),
	}
	block.SetHash()
	return &block
}

//3.生成哈希
func (block *Block) SetHash() {
	// 1.拼接数据
	blockInfo := append(block.PrevHash, block.Data...) // ... 第二个用法是slice可以被打散进行传递
	// 2.sha256
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]

}

//4.引入区块链
type BlockChain struct {
	// 定义一个区块链数组
	blocks []*Block
}

// 5.定义一个区块链
func NewBlockChain() *BlockChain{
	// 创建一个创世块，并作为第一个区块添加到区块链中
	genesisBlock := GenesisBlock()
	return &BlockChain{
		blocks: []*Block{genesisBlock},
	}
}

// 创世块
func GenesisBlock() *Block{
	return NewBlock("创世块", []byte{})
}

//6.添加区块
func (bc *BlockChain)AddBlock(data string)  {

	// 获取最后一个区块
	lastBlock := bc.blocks[len(bc.blocks)-1]
	prevHash := lastBlock.Hash

	// 创建新的区块
	block := NewBlock(data, prevHash)

	// 添加到区块链数组中
	bc.blocks = append(bc.blocks, block)
}


//6.重构代码

func main() {

	bc := NewBlockChain()
	bc.AddBlock("第二个块")
	bc.AddBlock("第三个块")
	for i, block := range bc.blocks{
		fmt.Printf("========= 当前区块高度 %d =========\r\n", i)
		fmt.Printf("前区块哈希： %x\n", block.PrevHash)
		fmt.Printf("当前区块哈希： %x\n", block.Hash)
		fmt.Printf("区块数据： %s\n", block.Data)
	}
}

