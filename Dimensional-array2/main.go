package main

import "fmt"

func main() {
	var num [3][3]int = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	for i := 0; i < len(num); i++ {
		for j := 0; j < len(num[i]); j++ {
			fmt.Printf("%v\t", num[i][j])
		}
		fmt.Println("")
	}
}
