// Contador de palabras
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func contarPalabras(frase string) int {
	palabras := strings.Fields(frase)
	return len(palabras)
}

func main() {
	fmt.Println("Ingrese una frase:")
	reader := bufio.NewReader(os.Stdin)
	frase, _ := reader.ReadString('\n')

	numPalabras := contarPalabras(frase)
	fmt.Println("Número de palabras:", numPalabras)
}
