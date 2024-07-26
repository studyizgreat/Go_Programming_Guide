package main

import "fmt"

func calc(num1 int, num2 int) (int, int) {
	addition := num1 + num2
	substraction := num1 - num2
	return addition, substraction
}
func main() {
	add_total, _ := calc(10, 20)
	_, sub_total := calc(10, 20)

	fmt.Println("Addition result: ", add_total)
	fmt.Println("Substraction result: ", sub_total)
}
