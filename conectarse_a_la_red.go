package main

import (
	"bufio"
	"fmt"
	"net"
)

var ips = make([]string, 0)

func connect_to_network(ip_to_connect string) {
	fmt.Println("Uniendose a la red...")
	con, _ := net.Dial("tcp", ip_to_connect)
	defer con.Close()
	r := bufio.NewReader(con)
	msg := "soy nuevo = ...."
	fmt.Println("Enviando mi ip...")
	fmt.Fprint(con, msg)
	fmt.Println("Recibiendo las demas ips...")
	resp, _ := r.ReadString('\n')
	fmt.Print(resp)
	// agregar aqui...
	fmt.Println("Listo.")
}

func main() {
	connect_to_network("10.11.97.218:8000")
}
