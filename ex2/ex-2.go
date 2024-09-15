package main2

import (
	"fmt"
	"reflect"
)

func main() {
	var a int = 2
	var b string = "hello"
	var c bool
	var f float64 = 6.223
	y := "This is string"

	fmt.Println(a, b, f, y)
	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(b), reflect.TypeOf(c), reflect.TypeOf(f), reflect.TypeOf(y))
}

// := used only inside of a func
