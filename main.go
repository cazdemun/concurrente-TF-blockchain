package main

import (
	"net"
	"fmt"
	"bufio"
	"encoding/json"
	"strconv"
	"time"
	"os"
	"strings"
	// "github.com/BurntSushi/toml"

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
		go requestHash(getBCPort(n), hashes)
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
	IPString := conf.Ip + ":" + conf.Port
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

func getBCPort(n string) string {
	s := strings.Split(n,":")
	
	nbcport,_ := strconv.Atoi(s[1])
	nbcports := strconv.FormatInt(int64(nbcport + 1), 10)

	return (s[0] + ":" + nbcports)
}

func startBlockchainServer(c Config) {
	ln, _ := net.Listen("tcp", c.Ip + ":" + c.Port)
	defer ln.Close()
	fmt.Printf("\n(Starting blockchain server: %s)\n\n", c.Ip + ":" + c.Port)
	
	for {
		con, _ := ln.Accept()
		go handle(con)
	}	
}


func sendNewIP(con net.Conn, msg Node) {
	defer con.Close()
	fmt.Fprintln(con, msg.toString())
}

func handleServerRequest(con net.Conn) {
	defer con.Close()

	r := bufio.NewReader(con)
	msg, _ := r.ReadString('\n')
	
	var tx Node
	json.Unmarshal([]byte(msg), &tx)

	switch tx.Instruction {
		case 0:			
			fmt.Println("Sending new node")
			for _, n := range conf.Neighbours {
				msg := 	Node{1, tx.Ip}
				con, _ := net.Dial("tcp", n)
				go sendNewIP(con, msg)
			}
			fmt.Fprintln(con, conf.toString())
			conf.append(tx.Ip)
			fmt.Println("Current neighbours ", conf.Neighbours)
		case 1:
			fmt.Println("Adding new node")
			conf.append(tx.Ip)
			fmt.Println("Current neighbours ", conf.Neighbours)
		default:
			fmt.Fprint(con, "Mensaje recibido, codigo incorrecto!")
	}

}

func joinNetwork() {
	if len(conf.Neighbours) == 0 {
		return
	}
	con, _ := net.Dial("tcp", conf.Neighbours[0])
	defer con.Close()

	msg := 	Node{0, conf.Ip + ":" + conf.DPort}
	fmt.Fprintln(con, msg.toString())
	
	r := bufio.NewReader(con)
	res, _ := r.ReadString('\n')
	fmt.Println("Respuesta: ", res)

	var dConf Config 
	json.Unmarshal([]byte(res), &dConf)

	for _, n := range dConf.Neighbours {
		conf.append(n)
	}
}

func startDiscoveryServer(c Config) {
	ln, _:= net.Listen("tcp", c.Ip + ":" + c.DPort)
	defer ln.Close()
	fmt.Printf("\n(Starting discovery server: %s) \n\n", c.Ip + ":" + c.DPort)
	for {
		con, _ := ln.Accept()
		go handleServerRequest(con)
	}
}
	
func genesisBC() BlockChain {
	var bc BlockChain
	
	block1 := Block{0, "1561451991", "", "0", "1@1@Charles Leon Urbano"}
	block1.Hash = block1.calculateHash()
	bc.append(block1)
	block2 := Block{1, "1561451992", "", block1.calculateHash(), "2@1@Charles Leon Urbano@Prognosis: Resfriado"}
	block2.Hash = block2.calculateHash()
	
	rand.Seed(time.Now().UTC().UnixNano())
	chance := rand.Intn(11)
	
	if chance < 5 {
		bc.append(block2)
	}
	
	return bc
}

var conf Config
var bc BlockChain

func main() {
	bc = genesisBC()
	fmt.Println(len(bc.Blockchain))
	conf.Ip = "localhost"

	// Input information

	pin := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your port: ")
	port,_ := pin.ReadString('\n')
	conf.DPort = strings.TrimRight(port, "\r\n")
	bcport,_ := strconv.Atoi(conf.DPort)
	conf.Port = strconv.FormatInt(int64(bcport + 1), 10)
	
	go startDiscoveryServer(conf)

	fmt.Print("Enter a neighbour IP and port: ")
	port,_ = pin.ReadString('\n')
	if strings.TrimRight(port, "\r\n") != conf.DPort {
		conf.append(strings.TrimRight(port, "\r\n"))
	}

	fmt.Println("Current neighbours ", conf.Neighbours)
	joinNetwork()
	// Starting Node Discovery


	// Starting blockchain

	go startBlockchainServer(conf)

	for {
		// Gathering Transaction Information
		gin := bufio.NewReader(os.Stdin)
		fmt.Print("Press enter to send transaction\n")
		info,_ := gin.ReadString('\n')
		
		// Building Transaction
		t := time.Now().Unix()
		msg := Message {1, info, strconv.FormatInt(t, 10)}
		newBlock := Block {bc.getLen(), msg.Timestamp, "", bc.LastBlock.Hash, msg.Payload}
		newBlock.Hash = newBlock.calculateHash()
		bc.append(newBlock)

		// Send Transaction
		for _, n := range conf.Neighbours  {
			con, _ := net.Dial("tcp", getBCPort(n))
			go sendTransaction(con, msg)
		}
	}
}