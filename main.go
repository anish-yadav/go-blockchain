package main

import (
	"fmt"
	"strconv"

	"github.com/anish-yadav/go-blockchain/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("Fist block after genesis")
	chain.AddBlock("Second block after genesis")
	chain.AddBlock("Third block after geensis")

	for _, block := range chain.Blocks {
		fmt.Printf("Prev hash:  %x\n", block.PrevHash)
		fmt.Printf("Data in block :  %s\n", block.Data)
		fmt.Printf("Hash:  %x\n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW : %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
