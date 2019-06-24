// package main

// import (
// 	"fmt"
// 	"net"
// )
// type Block struct {
// 	Index     int // 01
// 	Timestamp string // "2019-06-24 07:38:01.9108652 -0500 DST m=+0.034199001"
// 	Hash      string // "0x16b4434f2097eb905f0ab03a007ceb79"
// 	PrevHash  string // "0x49d6138cedb0e2c1f5cbf639ba7f5b68"
// 	Payload   string // "1@asd@asdas@asdas"
// }

// var Blockchain []Block

package main

import (
	"net"
	"fmt"
	"bufio"
)

func handle(con net.Conn) {
	defer con.Close()

	fmt.Println("Manejando mensaje")
	r := bufio.NewReader(con)

	for {
		msg, _ := r.ReadString('\n')
		fmt.Fprint(con, "test")
		fmt.Println("recibido: ", msg)
		
		//fmt.Println(con, msg)
		
		// if len(msg) == 0 ||  msg[0] == 'x' {
			// 	break
			// }
		}
	}
	
	func main() {
		ln, _ := net.Listen("tcp", "localhost:8000")
		defer ln.Close()
		
		fmt.Println("Esperando: ")
		con, _ := ln.Accept()
		defer con.Close()
		
		r := bufio.NewReader(con)
		msg, _ := r.ReadString('\n')
		fmt.Println("Recibido: ", msg)
		fmt.Fprint(con, msg)


	// for {
	// 	fmt.Println("k pasa aka")
	// 	con, _ := ln.Accept()
	// 	go handle(con)
	// }
}

// go frameworks servidor