package blockchain

import (
	"fmt"
	"strconv"
)

type BlockChain struct {
	Blocks []*Block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func (chain *BlockChain) Print() {
	for _, block := range chain.Blocks {

		fmt.Printf("Prev Hash: \t%x\n", block.PrevHash)
		fmt.Printf("Data: \t\t%s\n", block.Data)
		fmt.Printf("Hash: \t\t%x\n", block.Hash)
		pow := NewProof(block)
		fmt.Printf("Proof of Work: \t%s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println("---")
	}
}
