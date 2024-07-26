package main

import "fmt"

func main() {
	name_slice := []string{"Ali", "Erik", "James", "Ravi", "David"}

	for i := 0; i < len(name_slice); i++ {
		fmt.Println(name_slice[i])
	}
}
