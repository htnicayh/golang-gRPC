package main

import "fmt"

// func main() {
// 	var a int = 5
// 	var p *int = &a

// 	fmt.Println("Value stored in variable a = ", a)
// 	fmt.Println("Address of variable a = ", &a)
// 	fmt.Println("Value stored in variable p = ", int(*p))

// 	*p = 10

// 	fmt.Println("Value stored in variable a after = ", a)
// }

func main() {
	var a int = 5

	changeValue(a)
	fmt.Println(a)

	changeAddress(&a)
	fmt.Println(a)

}

func changeValue(a int) {
	a = 100
}

func changeAddress(a *int) {
	*a = 200
}
