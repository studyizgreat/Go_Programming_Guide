package main

import "fmt"

func main() {
	var a int = 42
	var p *int = &a // p stores the address of a

	var pointer2_add = &a
	fmt.Println(pointer2_add, " => ", *pointer2_add)

	fmt.Println("Value of a:", a)          //value 42
	fmt.Println("Address of a:", &a)       // Output: memory address of a
	fmt.Println("Pointer p:", p)           // Output: memory address of a
	fmt.Println("Value at pointer p:", *p) // Output: 42

	*p = 100                          // Modify the value at the address pointed by p
	fmt.Println("New value of a:", a) // Output: 100
}
