package main

import (
	"blockchain/blockchain"
	"fmt"
	"strconv"
)

func main() {
	chain := blockchain.InitBlockChain()
	chain.AddBlock("First Genesis")
	chain.AddBlock("Second Genesis")
	chain.AddBlock("Third Genesis")
	for _, block := range chain.Blocks {
		fmt.Printf("Previos Hash: %x\n", block.PrevHash)
		fmt.Printf("Data Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := blockchain.NewProof(block)
		fmt.Printf("pow %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

	}
}
