package main

import (
	"fmt"
	"sync"
)

func philosopher(name string, right, left sync.Mutex) {
	for {
		fmt.Println(name, "Pensando!")
		right.Lock()
		left.Lock()
		fmt.Println(name, "Comiendo!")
		right.Unlock()
		left.Unlock()
	}
}

func main() {
	fork := make([]sync.Mutex, 5)
	go philosopher("Socrates", fork[0], fork[1])
	go philosopher(" Aristoteles", fork[1], fork[2])
	go philosopher("  Nietzsche", fork[2], fork[3])
	go philosopher("   Platon", fork[3], fork[4])
	philosopher("    Fonsi", fork[0], fork[4])
}