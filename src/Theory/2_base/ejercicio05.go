// trivial concurrent code
package main

import (
	"fmt"
)

// variable compartida
var n int

func p() {
	k1 := 1
	n = k1
}
func q() {
	k2 := 2
	n = k2
}

func main() {
	//generar los procesos concurrentes
	go p() //goroutine
	q()

	//time.Sleep(time.Second * 2)
	fmt.Println("El resultado de n es: ", n)
}
