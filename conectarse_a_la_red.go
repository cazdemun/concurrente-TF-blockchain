package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
)

var routes = make([]Route, 0)

func connectToNetwork(sourceRoute Route, destinationRoute Route) {
	log.Println("Uniendose a la red...")

	con, err := net.Dial("tcp", fmt.Sprintf("%s:%s", destinationRoute.IP, destinationRoute.Port))
	checkError(err)
	defer con.Close()

	node, err := json.Marshal(Node{Route: sourceRoute, Instruction: 0})
	checkError(err)

	log.Println("Enviando datos del nodo...")
	fmt.Printf("Nodo: %v\n", string(node))
	fmt.Fprintln(con, string(node))

	log.Println("Recibiendo datos de los demás nodos...")
	r := bufio.NewReader(con)
	resp, err := r.ReadString('\n')
	checkError(err)
	log.Println("Rutas recibidas:", resp)
	json.Unmarshal([]byte(resp), &routes)
	routes = append(routes,destinationRoute)
	log.Println("Listo.")
}

func main() {

	// Pasar la ip de algun nodo por argumento del programa
	// ej: 'go run conectarse_a_la_red.go 192.168.1.50'
	destinationIP := os.Args[1]
	myRoute := Route{IP: getPrivateIP(), Port: PORT}
	log.Println("NODE IP: ", myRoute.IP)
	destinationRoute := Route{IP: destinationIP, Port: PORT2}
	connectToNetwork(myRoute, destinationRoute)

}
