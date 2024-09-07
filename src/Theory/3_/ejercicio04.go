// Tercer intento
package main

import (
	"fmt"
	"time"
)

var wantp bool = false
var wantq bool = false

func p() {
	for {
		fmt.Println("Line01-SNC P")
		fmt.Println("Line02-SNC P")
		wantp = true //cambio
		for wantq != false {
			//espera P
		}

		fmt.Println("Line01-SC P")
		fmt.Println("Line02-SC P")
		wantp = false
	}
}

func q() {
	for {
		fmt.Println("Line01-SNC Q")
		fmt.Println("Line02-SNC Q")
		wantq = true //cambio
		for wantp != false {
			//espera Q
		}

		fmt.Println("Line01-SC Q")
		fmt.Println("Line02-SC Q")
		wantq = false
	}
}

func main() {
	go p()
	go q()

	time.Sleep(time.Hour)
}
