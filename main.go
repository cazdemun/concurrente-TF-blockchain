// package main

// import (
// 	"fmt"
// 	"net"
// )
// type Block struct {
// 	Index     int // 01
// 	Timestamp string // "2019-06-24 07:38:01.9108652 -0500 DST m=+0.034199001"
// 	Hash      string // "0x16b4434f2097eb905f0ab03a007ceb79"
// 	PrevHash  string // "0x49d6138cedb0e2c1f5cbf639ba7f5b68"
// 	Payload   string // "1@asd@asdas@asdas"
// }

// var Blockchain []Block

package main

import (
	"net"
	"fmt"
	"bufio"

	"encoding/json"
	// "github.com/davecgh/go-spew/spew"
)

type Message struct {
	Type     int
	Payload string
}

func handleTx(con net.Conn, tx Message) {
	fmt.Fprint(con, "Mensaje recibido 1!")
	// 	hash := calculateHash(tx.message)
	//	var hashes []string // channel

	// 	for all IPs {
// 		go requestHash(IP, hashes [])
//	}

//	if (hash != majority(hashes)) {
	//		blockchain = requestLedger(random(IPs))
	//	}
}

func handleHash(con net.Conn) {
	fmt.Fprint(con, "Mensaje recibido 2!")
	//	sendHash(blockchain.last.hash)
}

func handleLedger(con net.Conn) {
	fmt.Fprint(con, "Mensaje recibido 3!")
//	sendLedger(blockchain.serialize)
}

func handle(con net.Conn) {
	defer con.Close()

	r := bufio.NewReader(con)
	msg, _ := r.ReadString('\n')
	
	var tx Message
	
	fmt.Println("Recibido: ", msg)
	
	json.Unmarshal([]byte(msg), &tx)
	//spew.Dump(tx)
	
	messageType := tx.Type
	
	if messageType == 1 {
		handleTx(con, tx)	
	} else if messageType == 2 {
		handleHash(con)
	} else if messageType == 2 {
		handleLedger(con)
	} else {
		fmt.Fprint(con, "Mensaje recibido!")
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
	
func main() {
	port := "8000"
	initBlockchainServer(port)
}