package main

import (
	"net"
	"fmt"
	"bufio"
	//"sync"
	// "time"
	// "strconv"
)

func send(con net.Conn, msg Message) {
	defer con.Close()
	
	fmt.Fprintln(con, msg.toString())
	
	r := bufio.NewReader(con)
	res, _ := r.ReadString('\n')
	fmt.Println("Respuesta: ", res)
}

func request(con net.Conn,msg Message) string {
	defer con.Close()
	
	fmt.Fprintln(con, msg.toString())
	
	r := bufio.NewReader(con)
	res, _ := r.ReadString('\n')
	return res
}

func sendTransaction(con net.Conn, msg Message) {
	// t := time.Now().Unix()
	// msg := Message {1, "1@doctor@algo@mas", strconv.FormatInt(t, 10)}
	send(con, msg)
}

//func requestHash(ip string, c chan string, wg *sync.WaitGroup) {
func requestHash(ip string, c chan<- string) {
	con, _ := net.Dial("tcp", ip)
	msg := Message {2, "", "0"}
	res := request(con, msg)
	c <- res
	//wg.Done()
	//fmt.Println("Respuesta: ", res)
}

// return ledger
func requestLedger(con net.Conn) {
	msg := Message {3, "", "0"}
	res := request(con, msg)
	fmt.Println("Respuesta: ", len(res))

}

// func main() {
// 	// gin := bufio.NewReader(os.Stdin)
	
// 	// for {
// 	// 	fmt.Print("Mensaje: ")
// 	// 	input, _ := gin.ReadString('\n')
// 	// 	fmt.Printf("Escribiste: %s", input)
// 	// }

// 	for {
// 		time.Sleep(2000 * time.Millisecond)
// 		con, _ := net.Dial("tcp", "localhost:8000")
// 		go sendTransaction(con)
// 		//go requestHash(con)
// 		//go requestLedger(con)
// 	}
// }