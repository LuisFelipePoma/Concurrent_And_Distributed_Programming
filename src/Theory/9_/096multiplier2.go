package main

import (
	"fmt"
)

var end chan bool

func zero(n int, west chan float64) {
	for i := 0; i < n; i++ {
		west <- 0.0
	}
	close(west)
}

func source(row []float64, south chan float64) {
	for _, element := range row {
		south <- element
	}
	close(south)
}

func sink(north chan float64) {
	for range north {
	}
}

func result(c [][]float64, i int, east chan float64) {
	j := 0
	for element := range east {
		c[i][j] = element
		j++
	}
	end <- true
}

func multiplier(firstElement float64, n int, north, east, south, west chan float64) {
	var sum, secondElement float64
	for i := 0; i < n; i++ {
		select {
		case secondElement = <-north:
			sum = <-east
		case sum = <-east:
			secondElement = <-north
		}
		sum = sum + firstElement*secondElement
		south <- secondElement
		west <- sum
	}
	close(south)
	close(west)
}

func main() {
	a := [][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	b := [][]float64{{1, 0, 2}, {0, 1, 2}, {1, 0, 0}}
	c := [][]float64{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	end = make(chan bool)
	nra := len(a)
	nca := len(a[0])
	ns := make([][]chan float64, nra+1) // canales norte sur, matrix de 4x3
	for i := range ns {
		ns[i] = make([]chan float64, nca)
		for j := range ns[i] {
			ns[i][j] = make(chan float64)
		}
	}
	ew := make([][]chan float64, nra) // canales easte oeste, matrix de 3x4
	for i := range ew {
		ew[i] = make([]chan float64, nca+1)
		for j := range ew[i] {
			ew[i][j] = make(chan float64)
		}
	}
	for i := 0; i < nra; i++ {
		go zero(nra, ew[i][nca])
		go result(c, i, ew[i][0])
	}
	for i := 0; i < nca; i++ {
		go source(b[i], ns[0][i])
		go sink(ns[nra][i])
	}
	for i := 0; i < nra; i++ {
		for j := 0; j < nca; j++ {
			go multiplier(a[i][j], nca,
				ns[i][j], ew[i][j+1],
				ns[i+1][j], ew[i][j])
		}
	}
	for i := 0; i < nra; i++ {
		<-end
	}
	fmt.Println(c)
}
