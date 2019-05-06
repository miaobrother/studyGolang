package main

import (
	"fmt"
	"time"
)

func processTask() {
	fmt.Println("Execute Something...")
}

func testFormat()  {
	now := time.Now()
	fmt.Println(now.Format("2006/01/02 15:04:05"))
}

func main() {
	ticker := time.Tick(time.Second)
	for i := range ticker {
		fmt.Printf("%v\n", i)
		processTask()
		testFormat()
	}

}
