package main

import (
	"fmt"
	"strings"
)

func main() {

	var sS = "The quick is a morning 日本"
	strSli := strings.Fields(sS)
	fmt.Printf("The type is%T\n",sS)
	fmt.Printf("The type is%v\n",strSli[0])
}
