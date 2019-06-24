package main

import (
	"net"
	"fmt"
	"bufio"
	// "os"
)

// func main() {
// 	con, err := net.Dial("tcp", "localhost:8000")
// 	if err != nil {
// 		fmt.Println("No se puede conectar.")
// 		return
// 	}
// 	defer con.Close()
// 	fmt.Fprintln(con, "Hola, soy cliente")
// }

func main() {
	// Dial, weird sintaxys for client
	con, _ := net.Dial("tcp", "localhost:8000")
	defer con.Close()
	fmt.Fprintln(con, "Hola, soy cliente")

	r := bufio.NewReader(con)
	msg, _ := r.ReadString('\n')
	fmt.Println("Respondido: ", msg)

	//gin := bufio.NewReader(os.Stdin)
	
	

	// for {
	// 	fmt.Print("\n Mensaje: ")
	// 	msg, _ := gin.ReadString('\n')
	// 	fmt.Fprint(con, msg)
	// 	resp, _ := r.ReadString('\n')
	// 	fmt.Print("Respuesta: ", resp)
	// 	// if msg[0] == 'x' {
	// 	// 	break
	// 	// }
	// }

	//fmt.Fprintln(con, "Hola, soy cliente")
}