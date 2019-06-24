package main

import (
	"log"
	"net"
	"os"
)

const PORT string = "8001"
const PORT2 string = "2005"

type Route struct {
	IP   string `json: "IP"`
	Port string `json:"Port"`
}

type Node struct {
	Route Route `json: "Route"`
	// 0 = nuevo, 1 = solo agrega
	Instruction int `json:"Instruction"`
}

func checkError(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

func getPrivateIP() string {
	addrs, err := net.InterfaceAddrs()
	checkError(err)
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
