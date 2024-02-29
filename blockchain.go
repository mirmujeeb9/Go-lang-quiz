package main

import (
	"fmt"
	"time"
)

type Block struct {
	Timestamp    int64
	Data         string
	PreviousHash string
	Hash         string
}


type Blockchain struct {
	blocks []*Block
}

func NewBlock(data string, previousHash string) *Block {
	return &Block{
		Timestamp:    time.Now().Unix(),
		Data:         data,
		PreviousHash: previousHash,
		
	}
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func (bc *Blockchain) DisplayAllBlocks() {
	for _, block := range bc.blocks {
		fmt.Printf("Timestamp: %d, Data: %s, PreviousHash: %s, Hash: %s\n",
			block.Timestamp, block.Data, block.PreviousHash, block.Hash)
	}
}

func (bc *Blockchain) ModifyBlock(index int, newData string) error {
	if index < 0 || index >= len(bc.blocks) {
		return fmt.Errorf("invalid block index")
	}
	bc.blocks[index].Data = newData
	
	return nil
}

func NewBlockchain() *Blockchain {
	genesisBlock := NewBlock("Genesis Block", "")
	return &Blockchain{
		blocks: []*Block{genesisBlock},
	}
}

func main() {
	
	bc := NewBlockchain()

	bc.AddBlock("First Block after Genesis")
	bc.AddBlock("Second Block after Genesis")
	bc.AddBlock("Third Block after Genesis")

	bc.DisplayAllBlocks()

	err := bc.ModifyBlock(1, "Modified Second Block")
	if err != nil {
		fmt.Println("Failed to modify block:", err)
		return
	}


	fmt.Println("\nAfter modification:")
	bc.DisplayAllBlocks()
}

