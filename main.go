package main

import "fmt"

func main() {
	var a int = 10
	//type assertion
	var f float64 = float64(a)

	fmt.Println(f)
}
