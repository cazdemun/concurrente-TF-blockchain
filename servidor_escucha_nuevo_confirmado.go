package main

import (
    "net"
    "fmt"
	"bufio"
)

var ips = make([]string,0)

func startServer() {
	fmt.Println("Iniciando Servidor guardador...")
	ln, _:= net.Listen("tcp", "10.11.97.218:8123")
	defer ln.Close()
	fmt.Println("Escuchando por puerto 8123")
	for {
		con, _ := ln.Accept()
		go handleServerRequest(con)
	}
}

func handleServerRequest(con net.Conn) {
	defer con.Close()
	r := bufio.NewReader(con)
	for {
		msg, _ := r.ReadString('\n')
		fmt.Print("IP nueva recibida: ",msg)
		ips = append(ips,msg)
	}
}

func main() {
}