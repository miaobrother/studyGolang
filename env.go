package main

import (
	"fmt"
	"os"
)

var (
	Home = os.Getenv("HOME")
	User = os.Getenv("USER")
	GOROOT = os.Getenv("GOROOT")
)

func main() {
	fmt.Println(Home,User,GOROOT)

}