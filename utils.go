package main

import (
	"encoding/json"
	"encoding/hex"
	"crypto/sha256"
	"strconv"
)

type Message struct {
	// 1: Send transaction with payload
	// 2: Request hash
	// 3: Request ledger
	Type     int
	Payload string
	Timestamp string
}

func (m Message) toString() string {
	stringByte, _ := json.Marshal(m)
	stringMessage := string(stringByte)
	return stringMessage
}

type Block struct {
	Index     int
	Timestamp string
	Hash      string
	PrevHash  string
	Payload   string
}

func (block Block) calculateHash() string {
	record := strconv.Itoa(block.Index) + block.Timestamp + block.PrevHash + block.Payload
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

type BlockChain struct {
	Blockchain []Block
	LastBlock Block
}

func (bc *BlockChain) append(b Block) {
	bc.Blockchain = append(bc.Blockchain, b)
	bc.LastBlock = b
}

func (bc BlockChain) toString() string {
	stringByte, _ := json.Marshal(bc)
	stringBC := string(stringByte)
	return stringBC
}

type Config struct {
  Port string
  Neighbours []string
}

func mostCommonHash(hashes []string) string {
	m := make(map[string]int)
  compare := 0
  var mostFrequent string

	for _, h := range hashes {
		word := h

		m[word] = m[word] + 1 

		if m[word] > compare { 
			 compare = m[word]  
			 mostFrequent = h
		}
	}
	
	return mostFrequent
}