package blockchain

import (
	"fmt"
	"strconv"
)

type BlockChain struct {
	Blocks []*Block
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
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
