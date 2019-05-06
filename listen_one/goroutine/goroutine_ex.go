package main

import (
	"fmt"
	"time"
)

func calc() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 1)
		fmt.Printf("The i is : %d\n", i)
	}
	fmt.Println("for execute done.")
}

func main() {
	go calc()
	fmt.Println(" main execute done...")
	time.Sleep(time.Second * 11)
}
