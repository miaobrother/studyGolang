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

func testSwitchV1()  {
	switch a:= 2;a {
	case 1:
		fmt.Printf("a=1\n")
	case 2:
		fmt.Printf("a=2\n")
	case 3:
		fmt.Printf("a=3\n")
	default:
		fmt.Printf("The default")
	}

}

func main() {
	testSwitch()
	testSwitchV1()
	//testNineNine()
}
