package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var utf8String = "This is a 新世界"
	fmt.Println(len([]rune(utf8String)))
	fmt.Println(utf8.RuneCountInString(utf8String))
}
