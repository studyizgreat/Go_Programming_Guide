// you cannot have an array value more then set

package main

import "fmt"

func main() {
	i := 0
	var num [5]int = [5]int{1, 2, 3, 4, 5}
	for i < len(num) {
		fmt.Printf("%v\t", num[i])
		i++
	}
	fmt.Printf("%v", num[5]) //this line will throw error

}
