// impresion de letras y números concurrente
package main

import (
	"fmt"
	"time"
)

func imprimeNumeros() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Número: ", i)
		time.Sleep(time.Millisecond * 100)
	}
}

func imprimeLetras() {
	for letra := 'a'; letra <= 'e'; letra++ {
		fmt.Println("Letra: ", string(letra))
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	go imprimeNumeros()
	go imprimeLetras()

	time.Sleep(time.Second * 3)
}
