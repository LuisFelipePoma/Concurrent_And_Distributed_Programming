// Números perfectos
package main

import "fmt"

func esNumeroPerfecto(num int) bool {
	sum := 0
	for i := 1; i < num; i++ {
		if num%i == 0 {
			sum += i
		}
	}
	return sum == num
}

func main() {
	fmt.Println("Números perfectos entre 1 y 10000:")
	for i := 1; i <= 10000; i++ {
		if esNumeroPerfecto(i) {
			fmt.Println(i)
		}
	}
}
