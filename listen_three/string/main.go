package main

import (
	"fmt"
)

func testString() {
	var str = "Hello World retail"
	fmt.Printf("The str[0] is %v\n", str[0])
	for index,val := range str{
		fmt.Printf("The index is %v, val is %c\n",index,val)
	}

	var byteSlice []byte
	fmt.Printf("The byteSlice is%v\n",byteSlice)
	byteSlice =[]byte(str)
	byteSlice[0]= 101
	fmt.Printf("The Now byteSlice is%c\n",byteSlice)

	var china rune = '中'
	fmt.Printf("The china is %c\n",china)

	str = "中"
	var runeSlice []rune
	runeSlice = []rune(str)
	fmt.Printf("The runeSlice is %d\n",len(runeSlice))
}

func main() {
	testString()
}
