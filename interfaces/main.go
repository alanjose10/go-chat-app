package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

func (s square) area() float64 {
	return s.side * s.side
}
func (s square) perimeter() float64 {
	return 4 * s.side
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

type shape interface {
	area() float64
	perimeter() float64
}

func printShapeInfo(s shape) {
	fmt.Printf("Area of %T is %f\n", s, s.area())
	fmt.Printf("Perimeter of %T is %.3f\n", s, s.perimeter())
}

func main() {
	s := square{side: 10}
	c := circle{radius: 5.43}
	printShapeInfo(s)
	printShapeInfo(c)
}
