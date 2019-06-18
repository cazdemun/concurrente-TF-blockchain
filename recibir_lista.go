package main
import (
	"fmt"
	"net"
	"bufio"
	"os"
)

var ips = make(int[])

func connect_to_network(ip_to_connect string) {
	fmt.Println("Uniendose a la red...")
	con,_ := net.Dial("tcp", ip_to_connect)
	defer con.Close()
	r := bufio.NewReader(con)
	for {
		fmt.Print("Mensaje: ")
		msg = "xd"
		fmt.Fprint(con,msg)
		resp,_ := r.ReadString('\n')
		fmt.Print("Respuesta: ",resp)
		if msg[0] == 'x' {
			break;
		}
	}
}

func main() {
	connect_to_network("10.11.97.218:8000")
}