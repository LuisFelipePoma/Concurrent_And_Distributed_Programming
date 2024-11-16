package main

import "fmt"

func main() {
	//chsync := make(chan int)
	chasync := make(chan int, 1)

	/*go func() {
		chsync <- 5
	}()
	fmt.Printf("%d\n", <-chsync)*/
	chasync <- 5
	fmt.Printf("%d\n", <-chasync)
}
