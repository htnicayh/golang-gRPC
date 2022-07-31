package main

import "fmt"

func main() {
	// Defer is used to ensure that a function call is performed
	// defer fmt.Println("Line 1")
	// fmt.Println("Line 2")
	// defer fmt.Println("Line 3")

	// Defer Stack
	throwAsNext()
	// numbers := 12
	// defer fmt.Println(numbers)

	// numbers = 20

	ten := 10
	zero := 0

	fmt.Println(ten / zero)

	/* Define panic: panic("Error") */
	// Panic is used to stop the program and print the error message
}

func throwAsNext() {
	fmt.Println("ThrowAsNext")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error: ", err)
		}
	}()
	fmt.Println("Done")
}
