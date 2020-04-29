package main

func main() {

	bc := NewBlockChain()
	bc.AddBlock("111111")
	bc.AddBlock("222222")

	//bc.AddBlock("第二个块")
	//bc.AddBlock("第三个块")
	//for i, block := range bc.blocks{
	//	fmt.Printf("========= 当前区块高度 %d =========\r\n", i)
	//	fmt.Printf("前区块哈希： %x\n", block.PrevHash)
	//	fmt.Printf("当前区块哈希： %x\n", block.Hash)
	//	fmt.Printf("区块数据： %s\n", block.Data)
	//}

}

