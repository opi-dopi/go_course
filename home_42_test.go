package main

import (
	"testing"
)

func TestSquareArea(t *testing.T) {
	var v float64
	var s Figure = Square{2}
	v = s.area()
	if v != 4 {
		t.Error("Expected 4, got ", v)
	}
}

func TestSquarePerimetr(t *testing.T) {
	var v float64
	var s Figure = Square{2}
	v = s.perimeter()
	if v != 8 {
		t.Error("Expected 8, got ", v)
	}
}

func TestCircleArea(t *testing.T) {
	var v float64
	var s Figure = Circle{2}
	v = s.area()
	if v != 12.566370614359172 {
		t.Error("Expected 12.566370614359172, got ", v)
	}
}

func TestCirclePerimetr(t *testing.T) {
	var v float64
	var s Figure = Circle{2}
	v = s.perimeter()
	if v != 12.566370614359172 {
		t.Error("Expected 12.566370614359172, got ", v)
	}
}
