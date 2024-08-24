// Palindromo
package main

import (
	"fmt"
	"strings"
)

func esPalindromo(palabra string) bool {
	palabra = strings.ToLower(palabra)
	for i := 0; i < len(palabra)/2; i++ {
		if palabra[i] != palabra[len(palabra)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var palabra string
	fmt.Println("Ingrese una palabra:")
	fmt.Scanln(&palabra)

	if esPalindromo(palabra) {
		fmt.Println("La palabra es un palíndromo")
	} else {
		fmt.Println("La palabra no es un palíndromo")
	}
}
