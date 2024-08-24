// Conversor de temperaturas
package main

import (
	// "fmt"
	// "time"
)

var n int
var k int = 10

func p() {
	for i := 0; i < k; i++ {
		print(i, " p\n")
		temp := n
		n = temp + 1
	}
}

func q() {
	for i := 0; i < k; i++ {
		print(i, " q\n")

		temp := n
		n = temp - 1
	}
}

func main() {
	go p()
	go q()
	// time.Sleep(time.Second)
	print(n, " ->\n")
}