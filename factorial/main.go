package main

import "fmt"

// Higher-order function that takes a function and applies it to a number
func apply(num int, operation func(int) int) int {
	return operation(num)
}

func main() {
	// Define some simple functions
	addOne := func(n int) int { return n + 1 }
	multiplyByTwo := func(n int) int { return n * 2 }

	// Use the higher-order function with different operations
	result1 := apply(5, addOne)        // Adds 1 to 5
	result2 := apply(5, multiplyByTwo) // Multiplies 5 by 2

	fmt.Println("5 after adding 1:", result1)
	fmt.Println("5 after multiplying by 2:", result2)
}
