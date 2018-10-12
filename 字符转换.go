package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a = 10
	fmt.Printf("The a type is %T\n",a)
	s := strconv.Itoa(a)
	fmt.Println(s)
}
