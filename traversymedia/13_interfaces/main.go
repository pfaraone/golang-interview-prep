package main

import (
	"fmt"
	"math"
)

// define out interface
type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

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
	fmt.Println("hello world")
	c := Circle{x: 0, y: 0, radius: 6}
	// r := Rectangle{width: 4, height: 6}

	fmt.Printf("Circle area: %f\n", c.area())
}
