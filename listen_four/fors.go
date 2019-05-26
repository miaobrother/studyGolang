package main

import (
	"fmt"
)

func testFor() {
	for i := 0; i < 100; i++ {
		fmt.Printf("The i is:%d\n", i)
	}
}

func testMultiSign() {
	for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 {
		fmt.Printf("%d * %d = %d\n", no, i, no*i)
	}
}

func testBreak() {
	for i := 0; i < 10; i++ {
		if i > 5 {
			break
		}
		fmt.Printf("The iii is:%d\n", i)
	}

}

func main() {
	testFor()
	testBreak()
	testMultiSign()
}
