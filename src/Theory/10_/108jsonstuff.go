package main

import (
	"encoding/json"
	"fmt"
)

// Alumno etc.
type Alumno struct {
	Codigo   string  `json:"code"`
	Nombre   string  `json:"name"`
	Promedio float32 `json:"grade"`
}

func main() {
	alumnos := []Alumno{
		{"1234ABC", "Renato Faltón", 12.49},
		{"3211XYZ", "Martín Ng", 12.49},
		{"cosasmm", "Juan José Mendoza", 12.49}}

	jsonBytes, _ := json.MarshalIndent(alumnos, "", "  ")
	jsonStr := string(jsonBytes)
	fmt.Println(jsonStr)

	var alumnos2 []Alumno
	json.Unmarshal(jsonBytes, &alumnos2)
	fmt.Println(alumnos2)
}
