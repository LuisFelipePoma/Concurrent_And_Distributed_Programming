package main

import "fmt"

func pipesort(id int, inCh, outCh chan int) {
	min := 100000000
	for num := range inCh {
		if num < min {
			outCh <- min
			min = num
		} else {
			outCh <- num
		}
	}
	fmt.Printf("Proceso %d: %d\n", id, min)
	close(outCh) // pon cerrar antes del printf y mira lo que ocurre, explique lo sucedido
}

func main() {
	nums := []int{6, 2, 4, 9, 1, 3, 8, 5, 10, 7}
	n := len(nums)
	ch := make([]chan int, n+1)
	for i := range ch {
		ch[i] = make(chan int)
	}
	for i := range nums {
		go pipesort(i, ch[i], ch[i+1])
	}
	go func() {
		for _, num := range nums {
			ch[0] <- num
		}
		close(ch[0])
	}()
	for range ch[n] {
	}
}
