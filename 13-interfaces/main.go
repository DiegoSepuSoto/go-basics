package main

import (
	"fmt"
	"math"
)

// Shape interface
type Shape interface {
	area() float64
}

// Circle struct
type Circle struct {
	x, y, radius float64
}

// Rectangle struct
type Rectangle struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	fmt.Println("Interfaces")

	circle := Circle{0, 0, 5}
	rectangle := Rectangle{10, 5}

	fmt.Println("Circle Area:", getArea(circle))
	fmt.Println("Rectangle Area:", getArea(rectangle))
}
