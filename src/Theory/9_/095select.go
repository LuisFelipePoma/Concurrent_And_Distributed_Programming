package main

import (
	"fmt"
	"math/rand"
	"time"
)

func proc(ch chan int) {
	for {
		dur := time.Duration(rand.Intn(50) + 100)
		time.Sleep(dur)
		ch <- 0
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go proc(ch1)
	go proc(ch2)
	for {
		select {
		case <-ch1: // valor leido es descartado, pero podríamos asignarlo a una variable y usarlo
			fmt.Println("Leido de canal 1 primero")
			<-ch2
		case <-ch2:
			fmt.Println("Leido de canal 2 primero")
			<-ch1
		}
	}
}
