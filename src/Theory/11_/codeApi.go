//creación de servicio Web - API

package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Alumno struct {
	Codigo  string `json:"cod"`
	Nombres string `json:"nom"`
	Dni     int    `json:"dni"`
}

var Alumnos []Alumno

func cargarDatos() {
	Alumnos = []Alumno{
		{"123", "Juan", 45678912},
		{"456", "Luis", 78945612},
		{"789", "Jorge", 1234578}}
}

func funcAlumnos(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	jsonBytes, _ := json.MarshalIndent(Alumnos, "", " ")
	io.WriteString(resp, string(jsonBytes))
	log.Println("Invoca listar alumnos")
}

func definirEndpoints() {
	http.HandleFunc("/Alumno/", funcAlumnos)

	log.Fatal(http.ListenAndServe(":9015", nil))
}

func main() {
	cargarDatos()
	definirEndpoints()
}
