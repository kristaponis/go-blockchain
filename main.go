package main

import (
	"fmt"
	"strconv"

	"github.com/kristaponis/go-blockchain/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("First block")
	chain.AddBlock("Second block")
	chain.AddBlock("Third block")

	for _, b := range chain.Blocks {
		fmt.Printf("Previous hash: %x\n", b.PrevHash)
		fmt.Printf("Data: %s\n", b.Data)
		fmt.Printf("Hash: %x\n", b.Hash)
		fmt.Println()

		pow := blockchain.NewProof(b)
		fmt.Println(strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
