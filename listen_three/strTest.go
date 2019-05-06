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

func reverse_Str() {
	var str = "hello world"
	var bytes = []byte(str)
	fmt.Printf("The bytes is:%v\n", string(bytes))
	for i:= 0;i < len(bytes)/2;i++{
		tmp_Bytes := bytes[len(str)-i-1]
		bytes[len(str)-i-1] = bytes[i]
		bytes[i] = tmp_Bytes
	}
}

func main() {
	//stringTest()
	reverse_Str()
}
