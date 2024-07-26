package main

import "fmt"

func main() {

	name_age := map[string]int{"Ali": 50, "Erik": 25, "Ravi": 33, "Sam": 44}

	fmt.Println(name_age) // printing out the map
	fmt.Println("The length of the map is: ", len(name_age))

	score_map := make(map[string]int)
	fmt.Println(score_map) //this is empty key map

	//accessing individual items in a map
	fmt.Printf("%v\t,%v\t", name_age["Ali"], name_age["Erik"])

	//adding a new key / value to the list
	name_age["Rahul"] = 24
	fmt.Println(name_age)
	// Rahul was added noted the values are not indexed or ordered as in array

	//updating a person age
	name_age["Ali"] = 40
	fmt.Println(name_age)

	//to delete the value you can use delete function
	delete(name_age, "Rahul")
	fmt.Println(name_age)

	fmt.Println("Looping over the list")
	// to loop over the list of key values
	for key, value := range name_age {
		fmt.Println(key, " => ", value)
	}

	fmt.Println("Deleting keys with iteration")
	// to loop over the list of key values
	for key, value := range name_age {
		fmt.Println(key, "=>", value)
		delete(name_age, key)
	}

	// OR you can simply re-initialize as below.
	// name_age = make(map[string]int)
	fmt.Println("Printing the new length of Map: ", len(name_age))

}
