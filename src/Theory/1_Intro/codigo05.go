// Fibonacci
package main

import "fmt"

func fibonacci(n int) []int {
	secuencia := make([]int, n)
	secuencia[0], secuencia[1] = 0, 1
	for i := 2; i < n; i++ {
		secuencia[i] = secuencia[i-1] + secuencia[i-2]
	}
	return secuencia
}

func main() {
	fmt.Println("Secuencia de Fibonacci (10 primeros elementos):", fibonacci(10))
}
