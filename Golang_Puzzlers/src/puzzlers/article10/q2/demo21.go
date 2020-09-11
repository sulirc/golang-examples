// package main

// import "fmt"

// func main() {
// 	// 示例1。
// 	ch1 := make(chan int, 1)
// 	ch1 <- 1
// 	// ch1 <- 2 // 通道已满，因此这里会造成阻塞。
// 	fmt.Println("after ch1")
// 	// 示例2。
// 	ch2 := make(chan int, 1)
// 	// elem, ok := <-ch2 // 通道已空，因此这里会造成阻塞
// 	// _, _ = elem, ok
// 	fmt.Println("after ch2")
// 	ch2 <- 1

// 	// 示例3。
// 	var ch3 chan int
// 	//ch3 <- 1 // 通道的值为nil，因此这里会造成永久的阻塞！
// 	//<-ch3 // 通道的值为nil，因此这里会造成永久的阻塞！
// 	_ = ch3
// }

package main

import (
	"fmt"
	"time"
)

func say(s string) {
	// for i := 0; i < 5; i++ {
	// 	time.Sleep(100 * time.Millisecond)
	// 	fmt.Println(s)
	// }
	fmt.Println("deadlock ch1 init", s)

	// deadlock example
	// ch1 := make(chan int, 1)
	// ch1 <- 1
	// ch1 <- 2

	// ch2 := make(chan int, 1)
	// ch2 <- 100
	// elem, ok := <-ch2
	// _, _ = elem, ok

	fmt.Println("after ch1", s)
}

func main() {
	go say("Hello World!")
	fmt.Println("Normal World!")
	time.Sleep(100 * time.Millisecond)
}
