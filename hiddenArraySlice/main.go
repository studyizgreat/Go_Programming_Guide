package main

import "fmt"

func main() {
	//under lying array is taken care by compiler
	initialSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	sub_slice := initialSlice[1:5]
	//print initial slice
	fmt.Println("Initial Slice: ", initialSlice)
	fmt.Println("Sub Slice: ", sub_slice)

	sub_slice[3] = 900
	fmt.Println("Sub slice after value modification", sub_slice)
	fmt.Println("Initial Slice after value modification in sub_slice\n ", initialSlice)

}

// 	for i := 6; i < 15; i++ {
// 		initialSlice = append(initialSlice, i)

// 	}

// 	fmt.Println("Updated slice: ", initialSlice)

// }
