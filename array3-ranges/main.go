package main

import "fmt"

func main() {

	var students [5]string = [5]string{"Ali", "James", "Ravi", "David", "Erik"}

	for index, element := range students {
		fmt.Println(index, " => ", "element", index, element)
	}

}
