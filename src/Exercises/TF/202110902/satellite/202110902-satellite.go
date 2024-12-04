package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"
)

// Estructura Zentraedi
type Zentraedi struct {
	ID string  `json:"id"`
	X  float64 `json:"x"`
	Y  float64 `json:"y"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run satellite.go <puerto_mecha1> <puerto_mecha2> ...")
		os.Exit(1)
	}

	mechaPorts := os.Args[1:]

	rand.Seed(time.Now().UnixNano())

	zentID := 0

	for {
		// Generar coordenadas y ID aleatorios para el Zentraedi
		z := Zentraedi{
			ID: fmt.Sprintf("%d", zentID),
			X:  rand.Float64() * 100,
			Y:  rand.Float64() * 100,
		}
		zentID++

		// Elegir un mecha aleatorio para enviar la ubicación del Zentraedi
		index := rand.Intn(len(mechaPorts))
		port := mechaPorts[index]

		// Enviar información del Zentraedi
		message := map[string]interface{}{
			"zentraedi": z,
		}
		data, _ := json.Marshal(message)
		sendData(port, data)

		fmt.Printf("Satélite envió Zentraedi ID %s en (%.2f, %.2f) al Mecha en el puerto %s\n", z.ID, z.X, z.Y, port)

		time.Sleep(20 * time.Second) // Enviar un nuevo Zentraedi cada 10 segundos
	}
}

// Envía datos a un puerto específico
func sendData(port string, data []byte) {
	conn, err := net.Dial("tcp", "localhost:"+port)
	if err != nil {
		fmt.Println("Error al conectar al puerto", port, ":", err)
		return
	}
	defer conn.Close()
	conn.Write(data)
}
