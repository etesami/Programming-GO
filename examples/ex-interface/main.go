package main

import "fmt"

// Circle is a type that implements the Shape interface.
type Circle struct {
	Radius float64
}

// Rectangle is another type that implements the Shape interface.
type Rectangle struct {
	Width  float64
	Height float64
}

// Area calculates the area of the circle.
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Area calculates the area of the rectangle.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	// Creating objects of types Circle and Rectangle
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 4, Height: 6}

	// Calculating the area of the circle and rectangle separately
	fmt.Println(circle.Area())
	fmt.Println(rectangle.Area())

	// Using the interface to calculate the area of different shapes
	printArea(circle)
	printArea(rectangle)

}

// Shape is an example interface with a method signature.
type Shape interface {
	Area() float64
}

// printArea is a function that takes a Shape and prints its area.
func printArea(s Shape) {
	fmt.Printf("Area: %f\n", s.Area())
}
