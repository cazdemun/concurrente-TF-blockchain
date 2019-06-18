package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go Productor(1,ch)
	go Consumidor(2,ch)
	time.Sleep(time.Second);
	fmt.Print("Acabo.\n")
}