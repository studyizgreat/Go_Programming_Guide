package main

import "fmt"

func main() {

	z := [...]int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println("Array 'z': ")
	for i := 0; i < len(z); i++ {
		fmt.Printf("%v\t %v\t", i, z[i])
		fmt.Println()
	}
}
