package main

import "fmt"

// type Human struct {
// 	name  string
// 	age   int
// 	phone string
// }

// type Student struct {
// 	Human
// 	school string
// 	loan   float32
// }

// type Employee struct {
// 	Human
// 	company string
// 	money   float32
// }

// // Define Interface
// type FirstInterface interface {
// 	SayHi()
// 	Sing(lyrics string)
// }

// func (h Human) SayHi() {
// 	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
// }

// func (h Human) Sing(lyrics string) {
// 	fmt.Println("La la la la...", lyrics)
// }

// func (e Employee) SayHi() {
// 	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name, e.company, e.phone)
// }

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct {
}

func (cw *ConsoleWriter) Write(data []byte) (int, error) {
	message, err := fmt.Println(string(data))
	return message, err
}

// =================================

type Increment interface {
	Increment() int
}

type count int

func (ic *count) Increment() int {
	*ic++
	return int(*ic)
}

// =================================

func main() {
	// var w Writer = &ConsoleWriter{}
	// w.Write([]byte("Hello, World!"))

	myInt := count(0)
	var inc Increment = &myInt

	for i := 0; i < 5; i++ {
		fmt.Println(inc.Increment())
	}

	// mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	// // paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	// // sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	// tom := Employee{Human{"Tom", 36, "444-222-XXX"}, "Things Ltd.", 5000}

	// // define interface i
	// var i FirstInterface

	// //i can store Student
	// i = mike
	// fmt.Println("This is Mike, a Student:")
	// i.SayHi()
	// i.Sing("November rain")

	// //i can store Employee
	// i = tom
	// fmt.Println("This is Tom, an Employee:")
	// i.SayHi()
	// i.Sing("Born to be wild")

	// // slice of FirstInterface
	// fmt.Println("Let's use a slice of FirstInterface and see what happens")
	// x := make([]FirstInterface, 3)
	// // these three elements are different types but they all implemented interface Men
	// x[0], x[1], x[2] = paul, sam, mike

	// for _, value := range x {
	// 	value.SayHi()
	// }

}
