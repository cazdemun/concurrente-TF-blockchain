package main

import (
	"net"
	"fmt"
	"bufio"
	"encoding/json"
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
	send(con, msg)
}

func requestHash(ip string, c chan<- string) {
	con, _ := net.Dial("tcp", ip)
	msg := Message {2, "", "0"}
	res := request(con, msg)
	c <- res
}

func requestLedger(ip string) BlockChain {
	con, _ := net.Dial("tcp", ip)
	msg := Message {3, "", "0"}
	res := request(con, msg)
	fmt.Println("Respuesta request ledger: ", res)

	var newBc BlockChain
	json.Unmarshal([]byte(res), &newBc)
	return newBc
}