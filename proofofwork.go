package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

// 定义一个工作量证明的结构ProfOfWork


type ProofOfWork struct {
	// a.block
	block *Block
	// b.目标值
	// 一个非常大的数，它有丰富的方法：比较、赋值方法
	target *big.Int
}

// 提供创建POW的函数
// NewProofOfWork(参数)
func NewProofOfWork(block *Block) *ProofOfWork {
	pow := ProofOfWork{
		block: block,
	}

	// 指定的难度值，字符串类型，需要转换
	targetStr := "0000f00000000000000000000000000000000000000000000000000000000000"
	// 引入的辅助变量，目的是将上面的难度值转成big.int
	tmpInt := big.Int{}
	// 将难度值赋值给big.int，指定16进制的格式
	tmpInt.SetString(targetStr, 16)
	pow.target = &tmpInt
	return &pow
}

// 提供计算不断计算hasn的函数
func (pow *ProofOfWork) Run() ([]byte, uint64) {

	// 1.拼装数据（区块的数据，还有不断变换的随机数）
    var nonce uint64
    block := pow.block
    var hash [32]byte

	for  {
		// 1.拼装数据（区块的数据，还有不断变换的随机数）
		tmp := [][]byte{
			Uint64ToByte(block.Version),
			block.PrevHash,
			block.MerkelRoot,
			Uint64ToByte(block.TimeStamp),
			Uint64ToByte(block.Difficulty),
			Uint64ToByte(nonce),
			block.Data,
		}
		blockInfo := bytes.Join(tmp, []byte{})
		// 2.做哈希运算
		hash = sha256.Sum256(blockInfo)
		// 3.与pow中的target进行比较
		tmpInt := big.Int{}
		// 将我们得到的hash数组转换成一个big.int
		tmpInt.SetBytes(hash[:])
		// 比较当前的哈希与目标哈希值，如果当前的哈希值小于目标的哈希值，就说明找到了
		//func (x *Int) Cmp(y *Int) (r int) {
		//	// x cmp y == x cmp y
		//	// x cmp (-y) == x
		//	// (-x) cmp y == y
		//	// (-x) cmp (-y) == -(x cmp y)
		if tmpInt.Cmp(pow.target) == -1{
			// a.找到了，退出返回
			fmt.Printf("挖矿成功！hash: %x\n, noce: %d\n", hash, nonce)
			return hash[:], nonce
		}else {
			// b.没找到，继续找，随机数加1
			nonce ++
		}
	}

	// return hash[:], nonce
}

//提供一个校验函数
//IsValid()