package main

import "fmt"

func testSwitch() {
	var a = "Hello"
	switch a {
	case "123", "345", "Hello":
		fmt.Printf("The a is %v\n", a)
		fallthrough
	default:
		fmt.Printf("The default a  is %v\n", a)

	}
	fmt.Printf("The a is %v\n", a)
}

func testNineNine() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d =%d\t", i, j, i*j)
		}
		fmt.Println()
	}
}

func main() {
	testSwitch()
	testNineNine()
}
