package main

import "fmt"

func main() {
	// Original slice
	src := []int{1, 2, 3, 4, 5}
	fmt.Println("Source Slice:", src)

	// Create a destination slice with smaller length
	dst := make([]int, 3)

	// Copy data from src to dst
	copied := copy(dst, src)
	fmt.Printf("Number of elements copied: %d\n", copied)

	// Print both slices
	fmt.Println("Destination Slice:", dst)
}
