package main

import "fmt"

// Define the Person struct
type Person struct {
	Age  int
	Name string
}

// Function that takes a struct by value
func PrintPerson(p Person) {
	// Modify the copy of the struct
	p.Age = 99
	fmt.Printf("Inside PrintPerson: Name: %s, Age: %d\n", p.Name, p.Age)
}

func main() {
	// Create an instance of Person
	p := Person{
		Age:  30,
		Name: "Alice",
	}

	// Call the function with the struct
	PrintPerson(p)

	// Print the original struct
	fmt.Printf("Outside PrintPerson: Name: %s, Age: %d\n", p.Name, p.Age)
}
