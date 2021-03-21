package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type BlockChain struct {
	blocks []*Block
}

type Block struct {
	PrevHash []byte
	Hash     []byte
	Data     []byte
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{prevHash, []byte{}, []byte(data)}
	block.DeriveHash()
	return block
}

func Genesis() *Block {
	return CreateBlock("GenesisBlock", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {
	chain := InitBlockChain()

	chain.AddBlock("First block")
	chain.AddBlock("Second block")
	chain.AddBlock("Third block")

	for _, b := range chain.blocks {
		fmt.Printf("Previous hash: %x\n", b.PrevHash)
		fmt.Printf("Data: %s\n", b.Data)
		fmt.Printf("Hash: %x\n", b.Hash)
		fmt.Printf("\n")
	}
}
