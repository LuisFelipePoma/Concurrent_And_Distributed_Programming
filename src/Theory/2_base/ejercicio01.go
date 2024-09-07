// sumar numeros pares
package main

import "fmt"

func sumarNumerosPares(n int) int {
	suma := 0
	for i := 0; i <= n; i += 2 {
		suma += i
	}

	return suma
}

func main() {
	fmt.Println("El resultado de la suma de pares es: ", sumarNumerosPares(20))
}
