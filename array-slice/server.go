package main

import "fmt"

func main() {
	points := [...]int{1, 2, 3}

	fmt.Printf("Before: %v, %T \n", points, points)

	var anyPoints [3]int
	// anyPoints[1] = 10

	fmt.Printf("After %v, %T \n", anyPoints, anyPoints)

	// Spread the points array into the anyPoints array
	spreadPoint := &points
	spreadPoint[0] = 10
	fmt.Printf("Spread: %v, %T \n", spreadPoint, spreadPoint)

	// Slices
	slices := [...]int{2, 5, 10, 23, 45, 67}

	a := slices[:]
	b := slices[2:4]
	c := slices[:6]
	d := slices[3:6]
	e := slices[0:4]

	// Change the value of the slice
	e[2] = 100

	fmt.Printf("Slices: %v, %T & Cap: %v \n", a, a, cap(a))
	fmt.Printf("Slices: %v, %T & Cap: %v \n", b, b, cap(b))
	fmt.Printf("Slices: %v, %T & Cap: %v \n", c, c, cap(c))
	fmt.Printf("Slices: %v, %T & Cap: %v \n", d, d, cap(d))
	fmt.Printf("Slices: %v, %T & Cap: %v \n", e, e, cap(e))

	// Using make function
	newSlices := make([]int, 3, 5)
	newSlices = append(newSlices, 1)
	fmt.Printf("New Slices: %v, %T & Cap: %v, Length: %v \n", newSlices, newSlices, cap(newSlices), len(newSlices))

	newSlices = append(newSlices, 2)
	fmt.Printf("New Slices: %v, %T & Cap: %v, Length: %v \n", newSlices, newSlices, cap(newSlices), len(newSlices))

	newSlices = append(newSlices, 3)
	fmt.Printf("New Slices: %v, %T & Cap: %v, Length: %v \n", newSlices, newSlices, cap(newSlices), len(newSlices))

	// Append different slice to slice
	// newSlices = append(newSlices, a...)
	newSlices = append(newSlices, []int{4, 5, 6}...)

	fmt.Printf("New Slices: %v, %T & Cap: %v, Length: %v \n", newSlices, newSlices, cap(newSlices), len(newSlices))
}
