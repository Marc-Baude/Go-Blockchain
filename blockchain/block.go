package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type BlockChain struct {
	Blocks []*Block
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

func (b *Block) CreateHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum224(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.CreateHash()
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
		fmt.Println("---")
	}
}
