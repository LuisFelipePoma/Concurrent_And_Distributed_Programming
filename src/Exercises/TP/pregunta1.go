package main

import (
	"fmt"
	"sync"
)

func main() {
	var data int
	var mu sync.Mutex  // Crear un Mutex para sincronizar

	go func() {
		mu.Lock()        // Bloquear antes de modificar
		data++
		mu.Unlock()      // Desbloquear después de modificar
	}()

	mu.Lock()            // Asegurarse que 'data' está correctamente actualizado antes de leer
	if data == 0 {
		fmt.Printf("The value is %v.\n", data)
	}
	mu.Unlock()
}
