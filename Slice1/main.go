package main

import "fmt"

func main() {
	//declare an array
	var arr [5]int = [5]int{1, 2, 3, 4, 5}

	//declare a slice
	slice1_arr_1 := arr[1:3] // this contains elements 2,3

	fmt.Println("Slice 1 array 1", slice1_arr_1)

	//to append
	slice1_arr_1 = append(slice1_arr_1, 7)
	fmt.Println("Slice 1 after append", slice1_arr_1)

	fmt.Println("Original array: ", arr)
}
