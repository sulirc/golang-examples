package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	fmt.Printf("hello, world\n")
	os.Exit(-1)
}
