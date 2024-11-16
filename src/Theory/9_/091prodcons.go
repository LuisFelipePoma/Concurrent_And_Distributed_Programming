package main

import (
	"fmt"
)

func producer(id int, ch chan string) {
	c := 0
	for {
		c++
		ch <- fmt.Sprintf("Producto %d producido por productor %d", c, id)
	}
}

func consumer(id int, ch chan string) {
	for {
		fmt.Printf("Consumidor %d consumiendo %s\n", id, <-ch)
	}
}

func main() {
	ch := make(chan string)
	for i := 0; i < 4; i++ {
		go producer(i, ch)
		go consumer(i, ch)
	}
	consumer(5, ch)
}
