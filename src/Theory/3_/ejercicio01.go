// Condición de carrera
package main

import (
	"fmt"
	"time"
)

func stingy(money *int) {
	for i := 0; i < 1000000; i++ {
		*money += 10
	}
}

func spendy(money *int) {
	for i := 0; i < 1000000; i++ {
		*money -= 10
	}
}

func main() {
	money := 100 //recurso compartido

	go stingy(&money)
	go spendy(&money)

	time.Sleep(time.Second * 10)

	fmt.Println("El saldo en la cuenta es: ", money)
}
