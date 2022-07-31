package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name     string
	age      int `Maximum can't over 100`
	address  string
	subjects []string
}

func newStruct() {
	var newPerson Person

	person := Person{
		name:    "Hyacinth",
		age:     1,
		address: "Jakarta",
		subjects: []string{
			"Math",
			"English",
			"Science",
		},
	}

	newPerson.name = "Hyacinth"
	newPerson.age = 1
	newPerson.address = "Jakarta"
	newPerson.subjects = []string{
		"Math",
		"English",
		"Science",
	}

	// Variables of Structs
	newStruct := struct{ name string }{name: "Hyacinth"}

	fmt.Printf("%v, %T\n", newStruct, newStruct)

	// Copy of Struct (New Memory)
	copyStruct := person
	copyStruct.name = "Amber"

	fmt.Printf("%v, %T\n", copyStruct, copyStruct)

	// Reflect
	var reflectPerson Person = Person{}
	ref := reflect.TypeOf(reflectPerson)
	field, _ := ref.FieldByName("age")

	fmt.Println(field)

	fmt.Println(newPerson)
	fmt.Println(person)
}

func main() {
	maps := map[string]int{
		"Hyacinth": 1,
		"Amber":    2,
		"Coral":    3,
	}

	// Reference Types (Values Type in Slices)
	newMaps := maps
	fmt.Printf("%v\n", newMaps)

	delete(newMaps, "Amber")

	// Default Value = 0 if Key not found in Maps
	_, isExist := newMaps["Amber"]
	fmt.Println(isExist)

	// make maps
	makeMaps := make(map[string]int)

	fmt.Printf("Before: %v, %T \n", maps, maps)
	fmt.Printf("After: %v, %T \n", makeMaps, makeMaps)

	newStruct()
}
