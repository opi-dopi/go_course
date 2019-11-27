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
	x, y float64
}

func (square Square) area() float64 {
	return square.x * square.y
}

func (square Square) perimeter() float64 {
	return square.x*2 + square.y*2
}

type Circle struct {
	r float64
}

func (circle Circle) area() float64 {
	return math.Floor(math.Pi*math.Pow(circle.r, 2)*100) / 100
}

func (circle Circle) perimeter() float64 {
	return math.Floor(2*math.Pi*circle.r*100) / 100
}

func main() {
	var s Figure = Square{2, 2}
	var c Figure = Circle{5}

	fmt.Println(s.area(), s.perimeter())
	fmt.Println(c.area(), c.perimeter())

}
