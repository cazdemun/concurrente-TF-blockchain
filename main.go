package main

import (
	"net"
	"fmt"
	"bufio"
	"encoding/json"
	"strconv"
	"time"
	"os"
	"github.com/BurntSushi/toml"

	"math/rand"
)

func handleTx(con net.Conn, msg Message) {
	// Adding new block
	newBlock := Block {bc.getLen(), msg.Timestamp, "", bc.LastBlock.Hash, msg.Payload}
	newBlock.Hash = newBlock.calculateHash()
	bc.append(newBlock)
	
	// Requesting hashes for verification
	hashes := make(chan string, len(conf.Neighbours))
	
	for _,n := range conf.Neighbours {
		go requestHash(n, hashes)
	}
	
	// Finding most common hash
	var hashesInfo []HashInfo
	
	for range conf.Neighbours {
		var hashInfo HashInfo
		json.Unmarshal([]byte(<-hashes), &hashInfo)
		hashesInfo = append(hashesInfo, hashInfo)

		fmt.Println("Hash from ", hashInfo.IP, ": ", hashInfo.Hash)
	}
	
	mostCommonHash := getMostCommonHash(hashesInfo)
	
	if newBlock.Hash != mostCommonHash {
		fmt.Println("Current hash does not match with majority")
		
		// Prunning IPs
		var selectablesIps []string
		for _,h := range hashesInfo {
			if h.Hash == mostCommonHash {
				selectablesIps = append(selectablesIps, h.IP)
			}
		}

		// Requesting and replacing current blockchain
		rand.Seed(time.Now().Unix())
		bc = requestLedger(selectablesIps[rand.Intn(len(selectablesIps))])
		
		fmt.Fprint(con, "houston we have problems")
	} else {	
		fmt.Fprint(con, "Transaccion recibida!")
	}
}
	
func handleHashRequest(con net.Conn) {
	IPString := "localhost:" + conf.Port
	hashInfo := HashInfo {bc.LastBlock.Hash, IPString}
	fmt.Fprint(con, hashInfo.toString())
}

func handleLedgerRequest(con net.Conn) {
	fmt.Fprint(con, bc.toString())
}

func handle(con net.Conn) {
	defer con.Close()
	
	r := bufio.NewReader(con)
	msg, _ := r.ReadString('\n')
	
	var tx Message
	json.Unmarshal([]byte(msg), &tx)
	fmt.Println("Peticion recibida: ", tx.Type)
	
	messageType := tx.Type
	
	switch messageType {
	case 1:
		handleTx(con, tx)	
	case 2:
		handleHashRequest(con)
	case 3:
		handleLedgerRequest(con)
	default:
		fmt.Fprint(con, "Mensaje recibido, codigo incorrecto!")
	}
}

func initBlockchainServer(port string) {
	ln, _ := net.Listen("tcp", "localhost:" + port)
	defer ln.Close()
	fmt.Printf("Opening blockchain server: %s \n\n", port)
	
	for {
		con, _ := ln.Accept()
		go handle(con)
	}	
}
	
func genesisBC() BlockChain {
	var bc BlockChain
	
	block1 := Block{1, "1561451991", "", "0", "asd@fgh"}
	block1.Hash = block1.calculateHash()
	bc.append(block1)
	block2 := Block{2, "1561451992", "", block1.calculateHash(), "asd@fgh"}
	block2.Hash = block2.calculateHash()
	bc.append(block2)
	block3 := Block{3, "1561451993", "", block2.calculateHash(), "asd@fgh"}
	block3.Hash = block3.calculateHash()
	bc.append(block3)
	
	rand.Seed(time.Now().UTC().UnixNano())
	chance := rand.Intn(11)
	
	if chance < 5 {
		bc.append(block3)
	}
	
	return bc
}

var conf Config
var bc BlockChain

func main() {
	bc = genesisBC()
	
	args := os.Args
	tomfile := args[len(args) - 1]
	fmt.Println(tomfile)

	toml.DecodeFile(tomfile, &conf)
	fmt.Println(len(bc.Blockchain))

	go initBlockchainServer(conf.Port)

	for {
		gin := bufio.NewReader(os.Stdin)
		fmt.Print("Press enter to send transaction\n")
		gin.ReadString('\n')
		
		t := time.Now().Unix()
		msg := Message {1, "1@doctor@algo@mas", strconv.FormatInt(t, 10)}
		newBlock := Block {bc.getLen(), msg.Timestamp, "", bc.LastBlock.Hash, msg.Payload}
		newBlock.Hash = newBlock.calculateHash()
		bc.append(newBlock)

		for _, n := range conf.Neighbours  {
			con, _ := net.Dial("tcp", n)
			go sendTransaction(con, msg)
		}
		
	}
}