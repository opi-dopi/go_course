package main

import (
	"fmt"
	"math"
)

type Figure interface {
	area() float64
	perimeter() float64
}

type Square struct {
	x float64
}

func (square Square) area() float64 {
	return square.x * square.x
}

func (square Square) perimeter() float64 {
	return square.x*2 + square.x*2
}

type Circle struct {
	r float64
}

func (circle Circle) area() float64 {
	return math.Pi*math.Pow(circle.r, 2)
}

func (circle Circle) perimeter() float64 {
	return 2*math.Pi*circle.r*100
}

func main() {
	var s Figure = Square{2}
	var c Figure = Circle{5}

	fmt.Println(s.area(), s.perimeter())
	fmt.Println(c.area(), c.perimeter())

}
