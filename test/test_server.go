package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func handle(con net.Conn) {
	defer con.Close()
	r := bufio.NewReader(con)
	fmt.Println("Conexion establecida, esperando instruccion...")
	msg, _ := r.ReadString('\n')
	fmt.Println("Recibido: ", msg)
	fmt.Fprintln(con, "buena!")
	fmt.Println("Respuesta enviada...")
	time.Sleep(time.Second)
	os.Exit(0)
	//sendMessage("bien", "192.168.1.50:8001", con)
}

func sendMessage(msg string, route string, con net.Conn) {
	//con, _ := net.Dial("tcp", route)
	//defer con.Close()
	fmt.Fprintln(con, msg)
	fmt.Println("Respuesta enviada...")
	os.Exit(0)
}

func main() {
	fmt.Println("Iniciando Servidor...")
	ln, _ := net.Listen("tcp", "192.168.1.50:8001")
	defer ln.Close()
	fmt.Println("Escuchando por puerto 8001")
	for {
		con, _ := ln.Accept()
		go handle(con)
	}
}
