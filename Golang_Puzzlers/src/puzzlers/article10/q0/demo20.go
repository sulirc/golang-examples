package main

import "fmt"

func main() {
	var firstNum = 2
	var copy = firstNum
	ch1 := make(chan int, 3)
	ch1 <- firstNum
	ch1 <- 1
	ch1 <- 3
	elem1 := <-ch1
	fmt.Printf("The first element received from channel ch1: %v\n",
		elem1)

	fmt.Printf("Address: firstNum %p & copy %p & elem1 %p\n", &firstNum, &copy, &elem1)

	<-ch1
	elemLast := <-ch1
	fmt.Printf("The last element received from channel ch1: %v\n",
		elemLast)
}
