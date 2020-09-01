package main

import (
	"fmt"
	"time"
)

func processTask() {
	fmt.Println("Execute Something...")
}

func testFormat() {
	now := time.Now()
	fmt.Println(now.Format("2006/01/02 15:04:05"))
}

func testTime() {
	now := time.Now()
	fmt.Printf("current time: %v\n", now)
	year := now.Year()
	month := now.Month()
	fmt.Printf("The year is %v, month is %v\n", year, month)
	timestamp := time.Now().Unix()
	fmt.Printf("The current timestamp is %v\n",timestamp)
}
func testTimestamp( timestamp int64)  {
	timeObject := time.Unix(timestamp,0)
	year := timeObject.Year()
	month := timeObject.Month()
	fmt.Printf("The year is %v, month is %v\n", year, month)
}

func testInterval()  {
	start := time.Now().UnixNano()
	for i := 0;i < 5; i++{
		time.Sleep(time.Millisecond)
		fmt.Printf("The i is %d\n",i)
	}
	end := time.Now().UnixNano()

	cost := end - start
	fmt.Printf("The cost is %v\n",cost/1000/1000)

}

func main() {
	//ticker := time.Tick(time.Second)
	//for i := range ticker {
	//	fmt.Printf("%v\n", i)
	//	processTask()
	//	testFormat()
	//}
	//testTime()
	//timestamps := time.Now().Unix()
	//testTimestamp(timestamps)
	testInterval()
}
