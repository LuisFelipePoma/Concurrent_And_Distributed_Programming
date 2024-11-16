package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numAdults      = 30 // Total de adultos mayores
	numTherapists  = 5  // Total de terapistas
	capacity       = 5  // Capacidad de la sala de rehabilitación
)

var (
	wg            sync.WaitGroup
	mutex         sync.Mutex
	adultsInRoom  = 0
)

func main() {
	wg.Add(numAdults) // Agregar el número total de adultos mayores al WaitGroup

	for i := 1; i <= numAdults; i++ {
		go adult(i) // Lanzar un goroutine para cada adulto mayor
	}

	wg.Wait() // Esperar a que todos los adultos mayores terminen
	fmt.Println("Todos los adultos mayores han terminado su rehabilitación.")
}

func adult(id int) {
	defer wg.Done() 

	mutex.Lock() // Adquirir el mutex para acceder a la sala
	for adultsInRoom >= capacity {
		mutex.Unlock() // Liberar el mutex si la sala está llena
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond) // Esperar antes de reintentar
		mutex.Lock() // Intentar nuevamente
	}
	adultsInRoom++ // 
	fmt.Printf("Adulto Mayor %d entra a la sala. Adultos en sala: %d/%d\n", id, adultsInRoom, capacity)
	mutex.Unlock() 

	// Simular la rehabilitación
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond) 

	mutex.Lock() // Adquirir el mutex para actualizar la sala
	adultsInRoom-- 
	fmt.Printf("Adulto Mayor %d sale de la sala. Adultos en sala: %d/%d\n", id, adultsInRoom, capacity) 
	mutex.Unlock() 

	// Simular la salida del adulto mayor
	fmt.Printf("Adulto Mayor %d ha completado su rehabilitación.\n", id)
}
