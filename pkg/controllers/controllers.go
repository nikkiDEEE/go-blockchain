package controllers

import (
	"crypto/sha256"
	"time"
	"encoding/hex"
	"encoding/json"
	"github.com/nikkiDEEE/go-blockchain/pkg/controllers"
	"github.com/nikkiDEEE/go-blockchain/pkg/models"
)

func newBlockchain() *models.Blockchain {
	return &models.Blockchain{[]*models.Block{GenesisBlock()}}
}

func genesisBlock() *models.Block {
	return CreateBlock(&models.Block{}, models.Event{IsGenesis: true})
}

func createBlock(prevBlock *models.Block, data interface{}) *models.Block {
	block := &models.Block
	block.PrevHash := prevBlock.PrevHash
	block.Position := prevBlock.Position + 1
	block.Data := data
	block.Timestamp := time.Now().String()
	block.generateHash()
}

func generateHash(block *models.Block) {
	bytes, _ := json.Marshal(block)
	data := block.PrevHash + string(block.Position) + string(bytes) + block.Timestamp
	hash := sha256.New()
	hash.Write([]byte(data))
	block.Hash = hex.EncodeToString(hash.Sum(nil))
}

func (bc *models.Blockchain) addBlock(data interface{}) {
	prevBlock := bc.blocks[len(bc.blocks) - 1]
	block := createBlock(prevBlock, data)
	if validBlock(prevBlock, block) {
		bc.blocks = append(bc.blocks, block)
	}
}

func validBlock(prevBlock, block *models.Block) bool {
	if prevBlock.Hash != block.PrevHash {
		return false
	}
	if !block.validateHash(block.Hash) {
		return false
	}
	if prevBlock.Position != (block.Position - 1) {
		return false
	}
	return true
}

func (block *models.Block) validateHash(hash string) {
	block.generateHash()
	if block.Hash != hash {
		return false
	}
	return true
}
