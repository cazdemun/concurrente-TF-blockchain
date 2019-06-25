// type Block struct {
// 	Index     int // 01
// 	Timestamp string // "2019-06-24 07:38:01.9108652 -0500 DST m=+0.034199001"
// 	Hash      string // "0x16b4434f2097eb905f0ab03a007ceb79"
// 	PrevHash  string // "0x49d6138cedb0e2c1f5cbf639ba7f5b68"
// 	Payload   string // "1@asd@asdas@asdas"
// }

package main

import (
	"net"
	"fmt"
	"bufio"
	"encoding/json"
	"strconv"
	"time"
	"os"
	//"sync"
	// "github.com/davecgh/go-spew/spew"
	"github.com/BurntSushi/toml"
)

func handleTx(con net.Conn, tx Message) {
	
	newBlock := Block {len(bc.Blockchain), tx.Timestamp, "", bc.LastBlock.Hash, tx.Payload}
	newBlock.Hash = newBlock.calculateHash()
	bc.append(newBlock)
	fmt.Println("own hash", newBlock.Hash)
	hashes := make(chan string, len(conf.Neighbours))
	
	//var wg sync.WaitGroup

	//wg.Add(len(conf.Neighbours))
	
	for _,n := range conf.Neighbours {
		// wg.Add(1)
		// go requestHash(n, hashes, &wg)
		go requestHash(n, hashes)
	}
	//wg.Wait()

	//close(hashes)
	
	//go func() {

		// for response := range hashes {
			// 		fmt.Println(response)
		// }
		var hashesS []string

		for range conf.Neighbours {
			// wg.Add(1)
			// go requestHash(n, hashes, &wg)
			// fmt.Println(<-hashes)
			hashesS = append(hashesS, <-hashes)
			fmt.Println(hashesS[len(hashesS) - 1])
		}
		fmt.Println("MOST COMMON")
		fmt.Println("MOST COMMON", mostCommonHash(hashesS))
		
	//}()
		

	//	if (hash != majority(hashes)) {
		//		blockchain = requestLedger(random(IPs))
		//	}
			
	// ll := len(bc.Blockchain)
	// fmt.Println("Block added: ", ll)
	// fmt.Fprint(con, "Transaccion recibida! - " + strconv.Itoa(ll))
	fmt.Fprint(con, "Transaccion recibida!")
}

func handleHash(con net.Conn) {
	fmt.Fprint(con, bc.LastBlock.Hash)
}

func handleLedger(con net.Conn) {
	fmt.Fprint(con, bc.toString())
}

func handle(con net.Conn) {
	defer con.Close()

	r := bufio.NewReader(con)
	msg, _ := r.ReadString('\n')
	
	var tx Message
	json.Unmarshal([]byte(msg), &tx)
	fmt.Println("Recibido Type: ", tx.Type)
	
	messageType := tx.Type

	switch messageType {
	case 1:
		handleTx(con, tx)	
	case 2:
		handleHash(con)
	case 3:
		handleLedger(con)
	default:
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
	
	return bc
}

var conf Config
var bc BlockChain
// var IPs = []string{"localhost:8000", "localhost:8001",	"localhost:8002", "localhost:8003"}

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
	fmt.Print("Press enter to send transaction")
	input, _ := gin.ReadString('\n')
	fmt.Println(input)
	
	t := time.Now().Unix()
	msg := Message {1, "1@doctor@algo@mas", strconv.FormatInt(t, 10)}
	newBlock := Block {len(bc.Blockchain), msg.Timestamp, "", bc.LastBlock.Hash, msg.Payload}
	newBlock.Hash = newBlock.calculateHash()
	bc.append(newBlock)

	fmt.Println("own:",newBlock.Hash)
		//time.Sleep(2000 * time.Millisecond)

		for _, n := range conf.Neighbours  {
			con, _ := net.Dial("tcp", n)
			go sendTransaction(con, msg)
		}
		//go requestHash(con)
		//go requestLedger(con)
	}
}