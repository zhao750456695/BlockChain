package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"log"
	"time"
)

//1.定义结构
type Block struct {
	// 1.版本号
	Version uint64

	// 2.前区块哈希
	PrevHash []byte

	// 3.Merkel根
	MerkelRoot []byte

	// 4.时间戳
	TimeStamp uint64

	// 5.难度值
	Difficulty uint64

	// 6.随机数，挖矿要找的数
	Nonce uint64


	// a.当前区块哈希
	Hash []byte
	// b.数据
	Data []byte

}

// 实现一个辅助函数，功能是将uint64转成[]byte
func Uint64ToByte(num uint64) []byte  {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil{
		log.Panic(err)
	}
	return buffer.Bytes()
}

//2.创建区块
func NewBlock(data string, prevBloackHash []byte) *Block {
	block := Block{
		Version: 00,
		PrevHash: prevBloackHash,
		MerkelRoot: []byte{},
		TimeStamp: uint64(time.Now().Unix()),
		Difficulty: 0,
		Nonce: 0,
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
	blockInfo = append(blockInfo, Uint64ToByte(block.Version)...)
	blockInfo = append(blockInfo, block.PrevHash...)
	blockInfo = append(blockInfo, block.MerkelRoot...)
	blockInfo = append(blockInfo, Uint64ToByte(block.TimeStamp)...)
	blockInfo = append(blockInfo, Uint64ToByte(block.Difficulty)...)
	blockInfo = append(blockInfo, Uint64ToByte(block.Nonce)...)
	blockInfo = append(blockInfo, block.Data...)

	// 2.sha256
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]

}