package main

import (
	"fmt"
)

func main() {
	var a = make(chan int, 1024)
	var b = make(chan int, 1024)

	for i := 0; i < 9; i++ {
		fmt.Printf("第%d次", i)
		a <- 1
		b <- 1
		select {
		case <-a:
			fmt.Println("from a")
		case <-b:
			fmt.Println("from b")
		}
	}
}
