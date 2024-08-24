// Ordenamiento de números
package main

import (
	"fmt"
	"sort"
	"container/list"
)

func ordenarNumeros(numeros []int) []int {
	sort.Ints(numeros)
	return numeros
}

func main() {
	listaNumeros := []int{4, 2, 8, 1, 6}
	// list
	test := list.New()
	test.Init()
	print(test.Back().Value)
	fmt.Println("Números ordenados:", ordenarNumeros(listaNumeros))
}
