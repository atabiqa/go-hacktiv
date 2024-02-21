package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	width, height float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}
func (r rectangle) area() float64 {
	return r.height * r.width
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}
func (r rectangle) perimeter() float64 {
	return 2 * (r.height + r.width)
}

func main() {
	var c1 = circle{radius: 5}
	var r1 = rectangle{width: 10, height: 5}

	//fmt.Printf("Type of c1 : %T", c1)
	//fmt.Printf("Type of r1 : %T", r1)

	fmt.Println("circle area:", c1.area())
	fmt.Println("circle perimeter:", c1.perimeter())

	fmt.Println("rectangle area:", r1.area())
	fmt.Println("rectangle perimeter:", r1.perimeter())
}