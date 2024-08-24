package main

import (
	"fmt"
	"math/rand"
	"time"
)

var n int

func p() {
	var temp int
	for i := 0; i < 10; i++ {
		temp = n
		pausita()
		n = temp + 1
		pausita()
	}
}

func main() {
	go p()
	go p()

	time.Sleep(time.Second)
	fmt.Println(n)
}

func pausita() {
	d := rand.Intn(50) + 50
	time.Sleep(time.Nanosecond * time.Duration(d))
}
