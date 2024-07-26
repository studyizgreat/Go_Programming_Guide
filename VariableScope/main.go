package main

import "fmt"

func main() {
	if true {
		// Block scope
		blockVar := "I am a block variable"
		// Accessible within the block
		fmt.Println(blockVar)
	}
	/**
	 fmt.Println(blockVar)
	 This would cause an error
	 because blockVar is not
	accessible here
	**/
}
