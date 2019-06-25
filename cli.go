package main

import (
	"net"
	"fmt"
	"bufio"
	
	"time"
	//"strconv"
	// "os"
)

func send(con net.Conn) {
	defer con.Close()

	fmt.Fprintln(con, "Hola, soy cliente ")
	
	r := bufio.NewReader(con)
	msg, _ := r.ReadString('\n')
	fmt.Println("Respuesta: ", msg)
}

func main() {
	// gin := bufio.NewReader(os.Stdin)
	
	// for {
		// 	fmt.Print("Mensaje: ")
		// 	input, _ := gin.ReadString('\n')
		// 	fmt.Printf("Escribiste: %s", input)
		// }

	for {
		time.Sleep(500 * time.Millisecond)
		con, _ := net.Dial("tcp", "localhost:8000")
		go send(con)
	}
}