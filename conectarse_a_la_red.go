package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

type Route struct {
	IP   string `json: "IP"`
	Port string `json:"Port"`
}

type Node struct {
	Route Route `json: "Route"`
	// 0 = nuevo, 1 = solo agrega
	Instruction int `json:"Instruction"`
}

var routes = make([]Route, 0)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getIP() string {
	addrs, err := net.InterfaceAddrs()
	check(err)
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	log.Fatal("No hay ip disponible")
	os.Exit(1)
	return ""
}

func connectToNetwork(routeToConnect string) {
	log.Println("Uniendose a la red...")
	con, err := net.Dial("tcp", routeToConnect)
	check(err)
	defer con.Close()
	node, err := json.Marshal(Node{Route: route, Instruction: 0})
	check(err)
	log.Println(string(node))
	log.Println("Enviando mi ip...")
	fmt.Fprintln(con, string(node))
	log.Println("Recibiendo las demas ips...")
	r := bufio.NewReader(con)
	time.Sleep(500 * time.Millisecond)
	resp, err := r.ReadString('\n')
	check(err)
	fmt.Println("Agregando las ips: ", resp)
	// TODO - agregar aqui...
	log.Println("Listo.")
}

var route Route

func main() {
	route = Route{IP: getIP(), Port: "8001"}
	log.Println("IP: ", route.IP)
	connectToNetwork("192.168.1.50:8001")
}
