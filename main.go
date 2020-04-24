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
//
//5.添加区块
//
//6.重构代码

func main() {

	block := NewBlock("a to b 1 coin", []byte{})

	fmt.Printf("前区块哈希： %x\n", block.PrevHash)
	fmt.Printf("当前区块哈希： %x\n", block.Hash)
	fmt.Printf("区块数据： %s\n", block.Data)

	fmt.Printf("hello btc")
}

