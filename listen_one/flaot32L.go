package listen_one

import (
	"fmt"
	"reflect"
)



func main() {
	var value1 float64
	value1 = 1
	value2 := 2
	value3 := 3.0

	v := value1 + value3
	fmt.Println(value1,value2,value3,v)

	fmt.Println(reflect.TypeOf(v)==reflect.ValueOf(v).Type())
	fmt.Println(reflect.ValueOf(v).Type())



}
