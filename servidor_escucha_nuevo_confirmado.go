package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
)

var activeRoutes = make([]Route, 0)

func startServer(sourceRoute Route) {
	fmt.Println("Iniciando Servidor guardador...")
	ln, err := net.Listen("tcp", fmt.Sprintf("%s:%s", sourceRoute.IP, sourceRoute.Port))
	checkError(err)
	defer ln.Close()
	fmt.Printf("Escuchando:\n IP %s - Puerto %s\n", sourceRoute.IP, sourceRoute.Port)
	for {
		con, err := ln.Accept()
		checkError(err)
		go handleServerRequest(con)
	}
}

func sendNewIP(route Route) error {
	con, err := net.Dial("tcp", fmt.Sprintf("%s:%s", route.IP, route.Port))
	if err != nil {
		return err
	}
	defer con.Close()

	routeRaw, err := json.Marshal(route)
	if err != nil {
		return err
	}

	fmt.Fprintln(con, string(routeRaw))

	return nil
}

func handleServerRequest(con net.Conn) {
	defer con.Close()
	r := bufio.NewReader(con)
	data, err := r.ReadString('\n')
	checkError(err)
	newNode := Node{}
	json.Unmarshal([]byte(data), &newNode)
	log.Printf("Nuevo nodo activo recibido %v\n", newNode)

	if newNode.Instruction == 0 {
		log.Println("Enviando la nueva ruta a los nodos activos...")
		for _, route := range activeRoutes {
			err := sendNewIP(route)
			checkError(err)
		}
	}

	log.Println("Enviando las rutas activas al nuevo nodo...")
	sendRoutes, err := json.Marshal(activeRoutes)
	fmt.Fprintln(con, string(sendRoutes))

	log.Println("Agregando nueva ruta a la lista...")
	activeRoutes = append(activeRoutes, newNode.Route)

	log.Println("Nueva ruta agregada.")
}

func main() {
	serverRoute := Route{IP: getPrivateIP(), Port: "2005"}
	startServer(serverRoute)
}
