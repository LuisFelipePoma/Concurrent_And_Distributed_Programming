// representar el algoritmo cena de los filósofos
package main

import (
	"fmt"
	"sync"
)

func filosofo(id int, tenedorIzd, tenedorDer sync.Mutex) {
	for {
		fmt.Printf("Filosofo %d, está pensando\n", id)
		tenedorIzd.Lock()
		tenedorDer.Lock()
		fmt.Printf("Filosofo %d, está comiendo\n", id)
		tenedorIzd.Unlock()
		tenedorDer.Unlock()
	}

}

func main() {
	tenedores := make([]sync.Mutex, 5)
	//4 procesos
	go filosofo(1, tenedores[0], tenedores[1])
	go filosofo(2, tenedores[1], tenedores[2])
	go filosofo(3, tenedores[2], tenedores[3])
	go filosofo(4, tenedores[3], tenedores[4])

	//proceso padre
	filosofo(5, tenedores[4], tenedores[0])
}
