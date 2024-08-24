// Números primos
package main

import "fmt"

func esPrimo(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var num int
	fmt.Println("Ingrese un número:")
	fmt.Scanln(&num)

	if esPrimo(num) {
		fmt.Println(num, "es primo")
	} else {
		fmt.Println(num, "no es primo")
	}
}
