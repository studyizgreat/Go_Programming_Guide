package main

import (
	"fmt"
)

func main() {
	fmt.Println("'Continue' For Loop")
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Printf("%v\t", i)
	}
	fmt.Println()
	fmt.Println("'Break' For Loop")
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Printf("%v\t", i)
	}
}
