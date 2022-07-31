package main

import "fmt"

var outside int = 10

// Export Global Variables
const Numbers int = 10

func main() {

	// Define Variables
	var numbers uint32
	numbers = 10
	var currentNumbers float32 = 13
	diffNumbers := 20

	var floatNumbers float32 = 3.2

	intNumbers := int(floatNumbers)

	fmt.Printf("%v %T \n", numbers, numbers)
	fmt.Printf("%v %T \n", currentNumbers, currentNumbers)
	fmt.Printf("%v %T \n", diffNumbers, diffNumbers)
	fmt.Printf("%v %T \n", floatNumbers, floatNumbers)
	fmt.Printf("%v %T \n", intNumbers, intNumbers)

	// Override Variables
	fmt.Printf("Before: %v %T \n", outside, outside)

	outside := 30

	fmt.Printf("After: %v %T \n", outside, outside)
}
