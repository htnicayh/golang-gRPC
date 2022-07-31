package main

import (
	"fmt"
)

const numbers int = 10
const (
	_ = iota
	one
	two
	three
)

func main() {
	fmt.Printf("Before: %v, %T \n", numbers, numbers)

	const numbers int = 20

	fmt.Printf("After %v, %T \n", numbers, numbers)
	fmt.Println(one, two, three)
}
