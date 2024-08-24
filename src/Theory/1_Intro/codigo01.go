// Suma de números pares.
package main

import "fmt"

func sumaNumerosPares(n int) int {
	suma := 0
	for i := 0; i <= n; i += 2 {
		suma += i
	}
	return suma
}

func main() {
	fmt.Println(sumaNumerosPares(10))
}
