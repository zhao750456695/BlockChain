package main

import (
	"fmt"
	"os"
)

// 接收命令行参数并且控制区块链操作
type CLI struct {
	bc *BlockChain
}

const Usage = `
    addBlock --data DATA    "add data to blockchain"
    printChain              "print all blockchain data"
`

// 接收参数
func (cli *CLI)Run()  {

	// block printChain
	// block addBlock
	// 1.得到所有的命令
	args := os.Args
	if len(args) < 2 {
		fmt.Printf(Usage)
		return
	}
	// 2.分析命令
	cmd := args[1]
	switch cmd {
	case "addBlock":
		// 添加区块
		fmt.Printf("添加区块")

		// 校验参数
		if len(args) == 4 && args[2] == "--data" {
			// a.获取数据
			data := args[3]
			// b.使用bc添加区块AddBlock
			cli.AddBlock(data)
		}else {
			fmt.Printf("添加区块参数使用不当，请检查！")
		}
	case "printChain":
		// 打印区块
		fmt.Printf("打印区块")
		cli.PrintBlockChain()
	default:
		fmt.Printf("无效的命令，请检查")
		fmt.Printf(Usage)
	}
	// 3.执行相应的命令
}


