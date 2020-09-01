package main

import (
	"fmt"
	"github.com/study/listen_two/access"
	"strings"
)

func testBool() {
	res := []string{"192.168.1.1", "192.168.1.2"}
	fmt.Println(strings.Join(res, "..."))
	var a bool
	fmt.Printf("The default a is %v\n", a)
	a = true
	fmt.Printf("The update after a is %v\n", a)
	a = !a
	fmt.Printf("Now a is %v\n", a)
	var b bool
	if a == true && b == true {
		fmt.Println("It's Ok")
	} else {
		fmt.Println("It's Bad")
	}

	if a == false || b == false {
		fmt.Println("|| data It's Ok")
	} else {
		fmt.Println("|| data It's Bad")
	}

}

func testInt() {
	var a int8
	a = 19
	fmt.Printf("The number a is: %d\n", a)
	a = -18
	fmt.Printf("The number a nos is: %d\n", a)
}

func testStr() {
	var a = "Hello World"
	fmt.Printf("The string a is: %s\n", a)
	c := "Default"
	fmt.Printf("The string c is: %s\n", c)
	c = "HHHH"
	fmt.Printf("The after string c is %v\n", c)
	clen := len(c)
	fmt.Printf("The clen is %v\n", clen)
	c = c + c
	fmt.Printf("The c + c is %v\n", c)
	c = fmt.Sprintf("%s%s", c, c)
	fmt.Printf("The sprintf c  is %v\n", c)
	ips := "192.168.1.1,192.168.1.2,192.168.1.3"
	ipSlice := strings.Split(ips, ",")
	fmt.Printf("The ipSlice is: %v\n", ipSlice)
	str := "www.baidu.com"
	index := strings.Index(str, "www")
	fmt.Printf("The index is:%v\n", index)
	var ipArray = []string{"10.222.1.1", "10.222.1.2"}
	res := strings.Join(ipArray, ";")
	fmt.Printf("The res is %v\n", res)
}

func testAccess() {
	fmt.Printf("The access is %v\n", access.A)

}
func main() {
	//testBool()
	//testInt()
	//testStr()
	testAccess()
}
