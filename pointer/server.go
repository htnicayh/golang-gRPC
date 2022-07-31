package main

import "fmt"

type newStruct struct {
	number int
}

func main() {
	// Create a pointer variable
	var numbers int = 12
	var seconds *int = &numbers

	fmt.Println(numbers, *seconds)

	numbers = 20
	fmt.Println(numbers, *seconds)

	arrayOne := []int{1, 2, 3}
	arrayTwo := arrayOne
	arrayOne[2] = 4

	fmt.Println(arrayOne, arrayTwo)

	// Pointer to a struct
	var newStructOne *newStruct
	newStructOne = new(newStruct)

	fmt.Println(newStructOne)

	// Pointer to a map
	var mapOne = map[string]string{"name": "John", "age": "30"}
	mapTwo := mapOne
	mapOne["name"] = "Jane"
	fmt.Println(mapOne, mapTwo)
}
