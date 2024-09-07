// input de datos / conversor de temperatura
package main

import "fmt"

func celciusToFahrenheit(celcius float64) float64 {
	return (celcius * 9 / 5) + 32
}
func fahrenheitToCelcius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}
func main() {
	var valTemp float64
	var unidad string
	//lectura de datos
	fmt.Print("Ingrese el valor de temperatura: ")
	fmt.Scanln(&valTemp) //captura por consola
	fmt.Print("\n" + "Ingrese la unidad (C/F): ")
	fmt.Scanln(&unidad) //captura por consola

	switch unidad {
	case "C":
		fmt.Println("El valor en Fahrenheit es: ", celciusToFahrenheit(valTemp))
	case "F":
		fmt.Println("El valor en Celcius es: ", fahrenheitToCelcius(valTemp))
	default:
		fmt.Println("Unidad no válida!")
	}

}
