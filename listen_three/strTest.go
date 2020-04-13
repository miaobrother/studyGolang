package main

import (
	"fmt"
)

func stringTest() {
	var international = "China"
	fmt.Printf("string[0]= %c\n The len is %d\n", international[0], len(international))
	for _, value := range international {
		fmt.Printf("The value is %c\n", value) //以字符显示 %c
	}
	var byteInter []byte //字符串本身不能被修改，如果要修改字符串，需要将字符串转换成 字节
	byteInter = []byte(international)
	for _, val := range byteInter {
		fmt.Printf("The byteInter value is %c\n", val)
	}
	byteInter[0] = 'G' // 必须使用单引号
	fmt.Printf("The byteInter is %c\n", byteInter)

	var chinaInter = "亚马逊称正在调整在中国的销售业务"
	var runeInter []rune
	runeInter = []rune(chinaInter)
	runeInter[0] = '你'
	fmt.Printf("The rune is %c\n", runeInter)

}

func reverseStr() {
	var str = "hello world"
	var bytes = []byte(str)
	fmt.Printf("The bytes is:%v\n", string(bytes))
	for i := 0; i < len(bytes)/2; i++ {
		tmpBytes := bytes[len(str)-i-1]
		bytes[len(str)-i-1] = bytes[i]
		bytes[i] = tmpBytes
	}
	str = string(bytes)
	fmt.Printf("The reverse results is %v\n", str)

	var str1 = "Hello"
	var bytes1 = []byte(str1)
	for i := 0; i < len(bytes1)/2; i++ {
		temp := bytes[len(str1)-i-1]
		bytes1[len(str1)-i-1] = bytes1[i]
		bytes1[i] = temp
	}
	str1 = string(bytes1)
	fmt.Printf("the bytes1 is %v\n", str1)

	var str3 = "World Wide"
	var bytes3 = []byte(str3)

	for i := 0; i < len(bytes3)/2; i++ {
		temp3 := bytes3[len(str3)-i-1]
		bytes3[len(str3)-i-1] = bytes3[i]
		bytes3[i] = temp3
	}
	str3 = string(bytes3)
	fmt.Printf("The bytes3 result: %v\n", str3)
}

func testRune() {
	var str = "上海自来水"
	var rune = []rune(str)
	for i := 0; i < len(rune)/2; i++ {
		temp := rune[len(rune)-i-1]
		rune[len(rune)-i-1] = rune[i]
		rune[i] = temp
	}
	str = string(rune)
	fmt.Printf("The rune is %v\n", str)
}

func testHuiwen() {
	var str = "上海自来水来自海上"
	var r = []rune(str)
	for i := 0; i < len(r)/2; i++ {
		temp := r[len(r)-i-1]
		r[len(r)-i-1] = r[i]
		r[i] = temp
	}
	str2 := string(r)
	fmt.Printf("The str2 is %v\n", str2)
	if str2 == str {
		fmt.Printf("The string str %v is Huiwen\n",str)
	} else {
		fmt.Printf("The string str %v not Huiwen\n", str)
	}
}

func main() {
	//stringTest()
	//reverseStr()
	//testRune()
	testHuiwen()
}
