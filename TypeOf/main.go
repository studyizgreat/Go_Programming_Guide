package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int = 10
	var b float64 = 20.9
	var c string = "A String"
	var d bool = true

	fmt.Printf("Value A:  %v\tType Of A is:  %T\n", a, reflect.TypeOf(a))
	fmt.Printf("Value B:  %v\tType Of B is:  %T\n", b, reflect.TypeOf(b))
	fmt.Printf("Value C:  %v\tType Of C is:  %T\n", c, reflect.TypeOf(c))
	fmt.Printf("Value D:  %v\tType Of D is:  %T\n", d, reflect.TypeOf(d))
}
