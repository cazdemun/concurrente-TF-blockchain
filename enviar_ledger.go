package main

import (
	// "crypto/sha256"
	// "encoding/hex"
	// "encoding/json"
	// "io"
	// "log"
	// "net/http"
	// "os"
	// "strconv"
	// "sync"
	"time"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"encoding/json"
	// "github.com/gorilla/mux"
	// "github.com/joho/godotenv"
)

// Block represents each 'item' in the blockchain
type Block struct {
	Index     int
	Timestamp string
	Hash      string
	PrevHash  string
	Payload   string
}

// var IPs []string = [
// 	"123123123",
// 	"123123123123",
// 	"1231231312312",
// 	"12312312312"
// ]

// Blockchain is a series of validated Blocks
var Blockchain []Block

func main() {
	var block1 Block = Block{
		1,
		"asd",
		"asdasd",
		"asdasd",
		"asdasd"}

	Blockchain = append(Blockchain, block1)
	
	var block2 Block = Block{
		2,
		"asd",
		"asdasd",
		"asdasd",
		"asdasd"}
	
	Blockchain = append(Blockchain, block2)
	
	var block3 Block = Block{
		3,
		"asd",
		"asdasd",
		"asdasd",
		"asdasd"}
		
	Blockchain = append(Blockchain, block3)
	
	spew.Dump(Blockchain)
	b, _ := json.Marshal(Blockchain)
	
	fmt.Println(string(b))
	
	var a string = string(b)

	var Blockchain2 []Block
	
	json.Unmarshal([]byte(a), &Blockchain2)
	Blockchain2 = append(Blockchain2, block3)
	spew.Dump(Blockchain2)

	t := time.Now()
	fmt.Println(t.String())
}