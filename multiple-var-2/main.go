package main

import (
	"fmt"
)

func main() {
	var (
		a int     = 10
		b float64 = 20.3
		c string  = "TestString"
	)

	fmt.Printf(
		"A: %v\n"+
			"B: %v\n"+
			"C: %v\n",
		a, b, c)
}
