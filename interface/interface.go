package main

import (
	"fmt"
)

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
	loan   float32
}

type Employee struct {
	Human
	company string
	salary  float32
}

type Man interface {
	SayHi()
	Song(lyrics string)
}

func (s *Student) SayHi() {
	fmt.Printf("Student Name is %v & phone %v, learning at %v school", s.name, s.phone, s.school)
}

func (h *Human) Song(lyrics string) {
	fmt.Printf("Hi, I am %v you can call me on %v, and my lyrics is %v", h.name, h.phone, lyrics)
}

func (e *Employee) SayHi() {
	fmt.Printf("Employee Name is %v & phone %v, working at %v company", e.name, e.phone, e.company)
}

func main() {

	fmt.Println("Hello, playground")
}
