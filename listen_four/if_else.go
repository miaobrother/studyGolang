package main

import "fmt"

func testIf() {
	num := 10
	if num%2 == 0 {
		fmt.Printf("The num is:%d\n", num)
	} else {
		fmt.Printf("The num is:%d\n", num)
	}
}

func main() {
	testIf()
}
