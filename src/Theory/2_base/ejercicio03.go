// Ordenamiento de numeros usando slice
package main

import (
	"fmt"
	"sort"
)

func ordenamiento(numeros []int) []int {
	sort.Ints(numeros)
	return numeros
}

func main() {
	lista := []int{10, 8, 1, 6, 4}
	fmt.Println("Numeros ordenados: ", ordenamiento(lista))
}
