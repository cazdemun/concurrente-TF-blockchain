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

func (bc BlockChain) getLen() int {
	return len(bc.Blockchain)
}

type Config struct {
  Port string
  Neighbours []string
}

type HashInfo struct {
  Hash string
  IP string
}

func (hash HashInfo) toString() string {
	stringByte, _ := json.Marshal(hash)
	stringHash := string(stringByte)
	return stringHash
}

func getMostCommonHash(hashes []HashInfo) string {
	m := make(map[string]int)
  compare := 0
  var mostFrequent string

	for _, h := range hashes {
		word := h.Hash

		m[word] = m[word] + 1 

		if m[word] > compare { 
			 compare = m[word]  
			 mostFrequent = h.Hash
		}
	}
	
	return mostFrequent
}