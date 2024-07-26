package main

import "fmt"

func main() {

	//var age int
	fmt.Println("Please enter your name and age")
	var name, age = fmt.Scanf("%s %d", &name, &age)
	fmt.Printf("Hi My name is %s and Age is %d", name, age)
}
