package main

import "fmt"

func (cli *CLI)PrintBlockChain()  {
	bc := cli.bc
	// 调用迭代器，返回我们的每一个区块数据
	it := bc.NewIterator()
	i := 0
	for  {
		// 返回区块，左移
		block := it.Next()
		fmt.Printf("========= 当前区块高度 %d =========\r\n", i)
		fmt.Printf("版本号： %d\n", block.Version)
		fmt.Printf("当前区块哈希： %x\n", block.Hash)
		fmt.Printf("梅克尔根： %x\n", block.MerkelRoot)
		fmt.Printf("时间戳： %d\n", block.TimeStamp)
		fmt.Printf("难度值： %d\n", block.Difficulty)
		fmt.Printf("随机数： %d\n", block.Nonce)
		fmt.Printf("前区块哈希： %x\n", block.PrevHash)
		fmt.Printf("区块数据： %s\n", block.Data)
		if len(block.PrevHash) == 0{
			fmt.Println("区块链遍历结束！")
			break
		}
		i ++
	}
}

func (cli *CLI)AddBlock(data string)  {
	cli.bc.AddBlock(data)
	fmt.Printf("添加区块成功！\n")
}