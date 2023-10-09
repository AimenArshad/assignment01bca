//Aimen-Arshad-20i-2375-Assignment-01

package assignment01bca

import (
	"crypto/sha256" // Importing the sha256 library for hashing
	"fmt"           // Importing the fmt library for printing
)

// Block represents a single unit in the blockchain.
type Block struct {
	Transaction  string // The transaction data stored in the block.
	Nonce        int    // Nonce is a random number used in the mining process.
	PreviousHash string // PreviousHash is the hash of the previous block.
	CurrentHash  string // CurrentHash is the hash of the current block.
}

// NewBlock creates and returns a new block with the provided data.
func NewBlock(transaction string, nonce int, previousHash string) *Block {
	block := &Block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
	}
	block.CurrentHash = CalculateHash(block)
	return block
}

// DisplayBlocks prints out the details of each block in the blockchain.
func DisplayBlocks(blocks []*Block) {
	for i, block := range blocks {
		fmt.Printf("Block %d:\n", i+1)
		fmt.Printf("Transaction: %s\n", block.Transaction)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Printf("Previous Hash: %s\n", block.PreviousHash)
		fmt.Printf("Current Hash: %s\n", block.CurrentHash)
		fmt.Println()
	}
}

// ChangeBlock updates the transaction data of a block and recalculates its hash.
func ChangeBlock(block *Block, newTransaction string) {
	block.Transaction = newTransaction
	block.CurrentHash = CalculateHash(block)
}

// VerifyChain checks if the blockchain is valid by verifying hashes and previous block references.
func VerifyChain(blocks []*Block) bool {
	for i := 1; i < len(blocks); i++ {
		currentBlock := blocks[i]
		previousBlock := blocks[i-1]

		if currentBlock.PreviousHash != previousBlock.CurrentHash {
			return false
		}

		if currentBlock.CurrentHash != CalculateHash(currentBlock) {
			return false
		}
	}
	return true
}

// CalculateHash computes the hash of a block based on its transaction data, nonce, and previous hash.
func CalculateHash(block *Block) string {
	data := fmt.Sprintf("%s%d%s", block.Transaction, block.Nonce, block.PreviousHash)
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}
