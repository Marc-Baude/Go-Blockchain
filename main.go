package main

import (
	"Go-Blockchain/blockchain"
	"fmt"
)

func main() {
	fmt.Println("My Block Chain v.0.0.1")

	chain := blockchain.InitBlockChain()
	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	chain.Print()
}
