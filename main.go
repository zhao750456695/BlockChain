package main

import "fmt"

func main() {

	bc := NewBlockChain()
	bc.AddBlock("111111")
	bc.AddBlock("222222")

    // 调用迭代器，返回我们的每一个区块数据
    it := bc.NewIterator()
    i := 0
	for  {
		// 返回区块，左移
		block := it.Next()
		fmt.Printf("========= 当前区块高度 %d =========\r\n", i)
		fmt.Printf("前区块哈希： %x\n", block.PrevHash)
		fmt.Printf("当前区块哈希： %x\n", block.Hash)
		fmt.Printf("区块数据： %s\n", block.Data)
		if len(block.PrevHash) == 0{
			fmt.Println("区块链遍历结束！")
			break
		}
		i ++
	}
}

