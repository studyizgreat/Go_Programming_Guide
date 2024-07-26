package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string = "50"
	var num int
	num, _ = strconv.Atoi(str)
	fmt.Printf("%d", num)
}
