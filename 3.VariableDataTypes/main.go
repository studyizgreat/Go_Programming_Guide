package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Variable")
	age := 23
	name := "Ali"
	var salary float64 = 200.15
	var favProgramLanguage string = "Go"

	fmt.Printf("%v is my name, I am %d years old, %f is my salary \n", name, age, salary)
	fmt.Printf("%v is my favorite Programming language\n", favProgramLanguage)

	fmt.Printf("Name variable is of type: %v\n", reflect.TypeOf(name))
	fmt.Printf("Age variable is of type: %v\n", reflect.TypeOf(age))
	fmt.Printf("Salary variable is of type: %v\n", reflect.TypeOf(salary))
	fmt.Printf("Favorite programming language variable is of type: %v\n", reflect.TypeOf(favProgramLanguage))

}
