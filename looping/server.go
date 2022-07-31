package main

import "fmt"

const (
	// Version is the current version of the program.
	Version = "0.0.1"
)

func main() {
	newArray := make([]int, 3)

	for i := 0; i < len(newArray); i++ {
		newArray[i] = i
	}

	// Diff Looping
	for index, value := range newArray {
		fmt.Println(index, value)
	}

	// Looping Maps
	newMaps := make(map[string]int)
	newMaps["Name"] = 1
	newMaps["Age"] = 2
	newMaps["Address"] = 3

	for key, value := range newMaps {
		fmt.Println(key, value)
	}

	// Looping Strings
	newString := "Hello World"
	for _, value := range newString {
		fmt.Printf(string(value))
	}

	fmt.Println()

	fmt.Println(newArray)
}
