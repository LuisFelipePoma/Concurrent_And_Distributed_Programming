// Conversor de temperaturas
package main

import "fmt"

func celsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

func fahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func main() {
	var temp float64
	var unidad string

	fmt.Println("Ingrese la temperatura:")
	fmt.Scanln(&temp)
	fmt.Println("Ingrese la unidad (C o F):")
	fmt.Scanln(&unidad)

	switch unidad {
	case "C":
		fmt.Println("Temperatura en Fahrenheit:", celsiusToFahrenheit(temp))
	case "F":
		fmt.Println("Temperatura en Celsius:", fahrenheitToCelsius(temp))
	default:
		fmt.Println("Unidad no válida")
	}
}
