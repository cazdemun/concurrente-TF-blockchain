package main

import (
	"net"
	"fmt"
	"bufio"
	
	"encoding/json"
	//"github.com/davecgh/go-spew/spew"
	"time"
	//"strconv"
	// "os"
)

type Message struct {
	// 1: Send transaction with payload
	// 2: Request hash
	// 3: Request ledger
	Type     int
	Payload string
}


func (m Message) toString() string {
	stringByte, _ := json.Marshal(m)
	stringMessage := string(stringByte)
	return stringMessage
}

func send(con net.Conn) {
	defer con.Close()
	
	firstMessage := Message {2, "1@doctor@algo@mas"}
	// stringFirstByte, _ := json.Marshal(firstMessage)
	// stringFirstMessage := string(stringFirstByte)
	
	//spew.Dump(firstMessage)
	//fmt.Println(stringFirstMessage)

	fmt.Fprintln(con, firstMessage.toString())
	
	r := bufio.NewReader(con)
	res, _ := r.ReadString('\n')
	fmt.Println("Respuesta: ", res)
}

func main() {
	// gin := bufio.NewReader(os.Stdin)
	
	// for {
		// 	fmt.Print("Mensaje: ")
		// 	input, _ := gin.ReadString('\n')
		// 	fmt.Printf("Escribiste: %s", input)
		// }


	for {
		time.Sleep(2000 * time.Millisecond)
		con, _ := net.Dial("tcp", "localhost:8000")
		go send(con)
	}
}