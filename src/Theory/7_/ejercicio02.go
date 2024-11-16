// semáforo condicional:  un recurso compartido y limite de procesos q acceden al RC
package main

import (
	"fmt"
	"sync"
	"time"
)

type semaforo struct {
	contador int
	limite   int
	mu       sync.Mutex
}

func (s *semaforo) wait() {
	s.mu.Lock()
	defer s.mu.Unlock()

	for s.contador >= s.limite {
		//espera activa del proceso
		s.mu.Unlock()
		//tiempo por hacer algo
		time.Sleep(time.Millisecond * 50)
		s.mu.Lock()
	}
	s.contador++
}

func (s *semaforo) signal() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.contador--
}

func proceso(id int, s *semaforo, wg *sync.WaitGroup) {
	defer wg.Done()
	s.wait()
	fmt.Printf("El proceso %d adquirió el semáforo\n", id)
	//simular trabajo del proceso
	time.Sleep(time.Millisecond * 50)
	fmt.Printf("El proceso %d liberó el semáforo\n", id)
	s.signal()
}

func main() {
	var wg sync.WaitGroup
	nroprocesos := 10
	limite := 3
	sem := &semaforo{limite: limite}

	for i := 0; i < nroprocesos; i++ {
		wg.Add(1)
		go proceso(i, sem, &wg)
	}

	wg.Wait()
	fmt.Printf("Todos los procesos culminaron!!!")
}
