package main

import "fmt"

type Point struct {
	X, Y float64
}

func (p *Point) Translate(dx, dy float64) {
	p.X = p.X + dx
	p.Y = p.Y + dy
}

func main() {
	var p Point = Point{3, 4}
	fmt.Println("Point p = ", p)

	p.Translate(1, 2)
	fmt.Println("Point p after translate = ", p)
}
