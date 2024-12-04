package main

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"net"
	"os"
	"sync"
	"time"
)

type Zentraedi struct {
	ID string  `json:"id"`
	X  float64 `json:"x"`
	Y  float64 `json:"y"`
}

type Info struct {
	MechaID  string  `json:"mecha_id"`
	Distance float64 `json:"distance"`
	ZentID   string  `json:"zent_id"`
}

type ZentraediInfo struct {
	Distances map[string]float64
	Done      chan struct{}
}

type Mecha struct {
	ID            string
	X             float64
	Y             float64
	Neighbors     []string
	ProcessedZent map[string]bool
	ZentraediMap  map[string]*ZentraediInfo
	Mutex         sync.Mutex
	TotalMechas   int
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Uso: go run mecha.go <mecha_id> <puerto> <puerto_vecino1> <puerto_vecino2> ...")
		os.Exit(1)
	}

	mechaID := os.Args[1]
	port := os.Args[2]
	neighborPorts := os.Args[3:]

	rand.Seed(time.Now().UnixNano())
	mecha := Mecha{
		ID:            mechaID,
		X:             rand.Float64() * 100,
		Y:             rand.Float64() * 100,
		Neighbors:     neighborPorts,
		ProcessedZent: make(map[string]bool),
		ZentraediMap:  make(map[string]*ZentraediInfo),
		TotalMechas:   4,
	}

	go startServer(&mecha, port)

	select {}
}

func startServer(mecha *Mecha, port string) {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
		return
	}
	fmt.Printf("---- Mecha %s escuchando en el puerto %s -----\n", mecha.ID, port)
	fmt.Printf("Ubicación: (%.2f, %.2f)\n", mecha.X, mecha.Y)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error de conexión:", err)
			continue
		}
		go handleConnection(mecha, conn)
	}
}

func handleConnection(mecha *Mecha, conn net.Conn) {
	defer conn.Close()
	decoder := json.NewDecoder(conn)
	var message map[string]interface{}
	err := decoder.Decode(&message)
	if err != nil {
		fmt.Println("Error al decodificar el mensaje:", err)
		return
	}

	if zentData, ok := message["zentraedi"]; ok {
		var z Zentraedi
		zData, _ := json.Marshal(zentData)
		json.Unmarshal(zData, &z)
		handleZentraedi(mecha, z)
	} else if infoData, ok := message["info"]; ok {
		var info Info
		infoBytes, _ := json.Marshal(infoData)
		json.Unmarshal(infoBytes, &info)
		handleInfo(mecha, info)
	}
}

func handleZentraedi(mecha *Mecha, z Zentraedi) {
	mecha.Mutex.Lock()
	if mecha.ProcessedZent[z.ID] {
		mecha.Mutex.Unlock()
		return
	}
	mecha.ProcessedZent[z.ID] = true
	// Inicializar la estructura para recopilar distancias
	mecha.ZentraediMap[z.ID] = &ZentraediInfo{
		Distances: make(map[string]float64),
		Done:      make(chan struct{}),
	}
	mecha.Mutex.Unlock()

	fmt.Printf("Mecha %s recibió Zentraedi ID %s en la posición (%.2f, %.2f)\n", mecha.ID, z.ID, z.X, z.Y)

	// Propagar información del Zentraedi a los vecinos
	message := map[string]interface{}{
		"zentraedi": z,
	}
	data, _ := json.Marshal(message)
	for _, neighborPort := range mecha.Neighbors {
		go sendData(neighborPort, data)
	}

	// Calcular distancia
	distance := calculateDistance(mecha.X, mecha.Y, z.X, z.Y)
	mecha.Mutex.Lock()
	mecha.ZentraediMap[z.ID].Distances[mecha.ID] = distance
	mecha.Mutex.Unlock()
	fmt.Printf("--- Mecha %s calculó distancia: %.2f ---\n", mecha.ID, distance)

	// Enviar distancia a los vecinos
	info := Info{
		MechaID:  mecha.ID,
		Distance: distance,
		ZentID:   z.ID,
	}
	infoMessage := map[string]interface{}{
		"info": info,
	}
	fmt.Printf("Mecha %s envía información a los mechas\n", mecha.ID)
	infoData, _ := json.Marshal(infoMessage)
	for _, neighborPort := range mecha.Neighbors {
		go sendData(neighborPort, infoData)
	}

	// Esperar hasta recibir todas las distancias
	go waitForDistances(mecha, z.ID)
}

func handleInfo(mecha *Mecha, info Info) {
	mecha.Mutex.Lock()
	defer mecha.Mutex.Unlock()

	zentID := info.ZentID
	if _, exists := mecha.ZentraediMap[zentID]; !exists {
		// Si el Zentraedi no está registrado, ignora la información
		return
	}

	mecha.ZentraediMap[zentID].Distances[info.MechaID] = info.Distance

	// Verificar si se han recibido todas las distancias
	if len(mecha.ZentraediMap[zentID].Distances) == mecha.TotalMechas {
		close(mecha.ZentraediMap[zentID].Done)
	}
}

func waitForDistances(mecha *Mecha, zID string) {
	mecha.Mutex.Lock()
	zInfo, exists := mecha.ZentraediMap[zID]
	mecha.Mutex.Unlock()

	if !exists {
		return
	}

	// Esperar a que todas las distancias sean recopiladas
	<-zInfo.Done

	mecha.Mutex.Lock()
	distance, existsMecha := zInfo.Distances[mecha.ID]
	if !existsMecha {
		fmt.Printf("Mecha %s no tiene distancia registrada para Zentraedi ID %s\n", mecha.ID, zID)
		fmt.Printf("--------------------------\n")
		mecha.Mutex.Unlock()
		return
	}
	minDistance := distance
	minMechaID := mecha.ID
	for id, dist := range zInfo.Distances {
		if dist < minDistance {
			minDistance = dist
			minMechaID = id
		}
	}
	// Resetear para el siguiente Zentraedi
	delete(mecha.ZentraediMap, zID)
	mecha.Mutex.Unlock()

	if mecha.ID == minMechaID {
		fmt.Printf("--------------------------\n")
		fmt.Printf("<- ¡Mecha %s lanza un cohete al Zentraedi ID %s! ->\n", mecha.ID, zID)
		fmt.Printf("--------------------------\n")
	}
}

func sendData(port string, data []byte) {
	conn, err := net.Dial("tcp", "localhost:"+port)
	if err != nil {
		fmt.Println("Error al conectar con el puerto", port, ":", err)
		return
	}
	defer conn.Close()
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("Error al enviar datos al puerto", port, ":", err)
	}
}

func calculateDistance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}
