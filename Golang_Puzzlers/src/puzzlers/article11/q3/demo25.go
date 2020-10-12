package main

import (
	"fmt"
	"time"
)

var channels = [3]chan int{
	nil,
	make(chan int),
	nil,
}

var myChan = make(chan int)

var numbers = []int{1, 2, 3}

func main() {
	// 执行顺序，从左到右，从上到下
	select {
	case getChan(0) <- getNumber(0):
		fmt.Println("The first candidate case is selected.")
	// 因为 make(chan int) 初始化的是不带缓冲的通道。非缓冲通道只有在收发双方都就绪的情况下才能传递元素值，否则就阻塞。
	case getChan(1) <- getNumber(1):
		fmt.Println("The second candidate case is selected.")
	case getChan(2) <- getNumber(2):
		fmt.Println("The third candidate case is selected")
	default:
		fmt.Println("No candidate case is selected!")
	}

	// v := <-channels[1]
	// fmt.Println("Get second channel value:", v)

	go testMyUnbufferedChan()
	myChan <- 100
	fmt.Println("Trigger after put value to myChan")

	time.Sleep(time.Second)
}

func testMyUnbufferedChan() {
	fmt.Println("Test myChan")

	v := <-myChan
	fmt.Println("Get value from myChan:", v)
}

func getNumber(i int) int {
	fmt.Printf("numbers[%d]\n", i)
	return numbers[i]
}

func getChan(i int) chan int {
	fmt.Printf("channels[%d]\n", i)
	return channels[i]
}
