package main

import "fmt"

func main() {

	arr1 := []int{1, 3, 5, 7, 9}
	arr2 := []int{2, 4, 6, 8, 10}

	arr_final := append(arr1, arr2...)
	fmt.Println(arr_final)

}
