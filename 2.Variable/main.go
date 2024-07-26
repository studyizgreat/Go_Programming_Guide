package main

import (
	"fmt"
)

func main() {
	fmt.Println("Variable")
	age := 23
	name := "Ali"
	var salary float64 = 200.15
	var favProgramLanguage string = "Go"

	fmt.Printf("%v is my name, I am %d years old, %f is my salary \n", name, age, salary)
	fmt.Printf("%v is my favorite Programming language\n", favProgramLanguage)
}
