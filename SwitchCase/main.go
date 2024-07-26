package main

import "fmt"

func main() {

	var score int = 100

	switch {
	case score >= 90:
		fmt.Println("Excellent!")
	case score >= 60:
		fmt.Println("Good")
	case score >= 40:
		fmt.Println("OK!")
	default:
		fmt.Println("Not ok!")
	}

}
