// primer intento para solución de probl. SC
package main

import (
	"fmt"
	"time"
)

var turno int = 1

func p() {
	for {
		fmt.Println("Line01-SNC P")
		fmt.Println("Line02-SNC P")
		//esperar que se cumpla la condición
		for turno != 1 {
			//espera P
		}
		fmt.Println("Line01-SC P")
		fmt.Println("Line02-SC P")

		turno = 2
	}
}

func q() {
	for {
		fmt.Println("Line01-SNC Q")
		fmt.Println("Line02-SNC Q")
		for turno != 2 {
			//espera Q
		}
		fmt.Println("Line01-SC Q")
		fmt.Println("Line02-SC Q")

		turno = 1
	}
}

func main() {
	go p()
	go q()

	time.Sleep(time.Hour)
}

//No cumple exclusion Mutua
// Line01-SC P
// Line01-SC Q
// Line02-SC P
