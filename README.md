# concurrente-TF-blockchain
Trabajo Final Programación Concurrente y Distribuida

Respetar la siguiente estructura...

```golang
type Route struct {
	IP string `json: "IP"`
	Port string `json:"Port"`
}

type Node struct {
	Route Route `json: "Route"`
	// 0 = nuevo, 1 = solo agrega
	Instruction int `json:"Instruction"`
}
```