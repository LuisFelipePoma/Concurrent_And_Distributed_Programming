// Recursividad: Factorial
package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func main() {
	fmt.Println("El factorial de 20 es ", factorial(20))
}
