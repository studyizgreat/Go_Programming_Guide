package main

import (
	"fmt"
)

// var fruit float64 = 20
// var fruit float64 = 9
var fruit float64 = 4.25

func main() {

	if fruit >= 10.50 {
		fmt.Println("Expensive")
	} else if fruit >= 5.50 && fruit <= 10.50 {
		fmt.Println("Normal")
	} else {
		fmt.Println("Cheap")
	}
}
