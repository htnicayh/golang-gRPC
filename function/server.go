package main

import "fmt"

const Default string = "Default"

func main() {
	// ...
	slices := []int{1, 2, 3, 4, 5}
	results := caculator(Default, slices)

	fmt.Println(results)

	// Struct
	person := newStruct{"John", 30}
	person.newPerson()
}

func caculator(strings string, a []int) (sum int) {
	for _, v := range a {
		sum += v
	}
	fmt.Println(strings)
	return
}

type newStruct struct {
	name string
	age  int
}

func (person newStruct) newPerson() {
	fmt.Println(person.name, person.age)
}
