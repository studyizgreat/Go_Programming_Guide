package main

import "fmt"

func main() {
	var fruit string                //declaring a variable
	fruit = "Apple"                 // assigning a value to the variable
	var vegetable string = "Carrot" // declare and initialize
	drink := "water"                // using short form implicit data type example
	price := 20.5                   // declaring an int variable
	// formatted print statement
	fmt.Printf("Fruit: %v\n"+
		"Vegetable: %v\n"+
		"Drink: %v\n"+
		"Price: %v\n",
		fruit, vegetable, drink, price)
}
