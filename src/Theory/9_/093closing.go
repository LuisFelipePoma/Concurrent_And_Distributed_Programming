package main

import "fmt"

func dostuff(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch) // channels pueden ser cerrados!
	// si no se cierra el canal, el for de abajo, causará deadlock
}

func main() {
	ch := make(chan int)
	// un canal sincrono como éste, no puede ser leido y escrito desde el mismo proceso
	// ch <- 5 // <- ésto por ejemplo dará error!
	go dostuff(ch)

	for num := range ch { // for each termina cuando canal se cierra
		fmt.Println(num)
	}
}
