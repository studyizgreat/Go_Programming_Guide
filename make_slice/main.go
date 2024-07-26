package main

import "fmt"

func main() {
	// Create a dynamic slice of 10 elements with values from 1 to 10
	dynamicSlice := make([]int, 10) // Create a slice of length 10
	for i := 0; i < len(dynamicSlice); i++ {
		dynamicSlice[i] = i + 1 // Fill the slice with values 1 to 10
	}

	// Print the original slice with index and values
	fmt.Println("Original Slice:")
	for index, value := range dynamicSlice {
		fmt.Printf("Index %d: Value %d\n", index, value)
	}

	// Create a sub-slice of 5 elements from the dynamic slice
	// Let's take elements from index 2 to index 6
	subSlice := dynamicSlice[2:7]

	// Print the sub-slice with index and values
	fmt.Println("\nSub-Slice:")
	for index, value := range subSlice {
		// Note: Index here is the index in the sub-slice, not in the original slice
		fmt.Printf("Sub-Slice Index %d: Value %d\n", index, value)
	}

	// Print the sub-slice starting index in the original slice
	fmt.Printf("\nSub-Slice starting index in the original slice: %d\n", 2)
}
