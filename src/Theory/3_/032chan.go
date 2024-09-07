package main

import (
	"fmt"
)

var ch chan bool

/*func p() {
	fmt.Println("Ping P")
	ch <- true
}*/

func main() {
	ch = make(chan bool)
	ch <- true
	//go p()
	<-ch
	fmt.Println("Pong P")
}
