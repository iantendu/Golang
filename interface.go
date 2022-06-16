package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}
type Circle struct {
	x, y, radius float64
}
type Rectangle struct {
	width, height float64
}

func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}
func (rec Rectangle) area() float64 {
	return rec.width * rec.height
}

func getArea(shape Shape) float64 {
	return shape.area()
}

func main() {

	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Rectangle{width: 6, height: 5}

	fmt.Printf("Circle area: %f\n", getArea(circle))
	fmt.Printf("Rectangle area: %f\n", getArea(rectangle))

}
