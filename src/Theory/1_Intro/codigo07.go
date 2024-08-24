// Conversor de número a texto
package main

import (
	"fmt"
	"strconv"
)

var unidades = []string{"", "uno", "dos", "tres", "cuatro", "cinco", "seis", "siete", "ocho", "nueve"}
var decenas = []string{"", "diez", "veinte", "treinta", "cuarenta", "cincuenta", "sesenta", "setenta", "ochenta", "noventa"}
var especiales = []string{"diez", "once", "doce", "trece", "catorce", "quince", "dieciséis", "diecisiete", "dieciocho", "diecinueve"}

func numeroEnPalabras(num int) string {
	if num < 10 {
		return unidades[num]
	} else if num < 20 {
		return especiales[num-10]
	} else if num < 100 {
		decena := decenas[num/10]
		unidad := unidades[num%10]
		if unidad == "" {
			return decena
		}
		return decena + " y " + unidad
	}
	return "Número fuera de rango"
}

func main() {
	var numStr string
	fmt.Println("Ingrese un número entero:")
	fmt.Scanln(&numStr)

	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Número en palabras:", numeroEnPalabras(num))
}
