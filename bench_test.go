package main

import (
	"fmt"
	"strconv"
	"testing"
)

var s string = "12345"
func BenchmarkSample(b *testing.B) {
	for i := 0; i < b.N; i++ {
		myStrToInt1(s)
	}
}

func BenchmarkSample1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		myStrToInt2(s)
	}
}

func myStrToInt1(s string) (int, error) {

	i, err := strconv.Atoi(s)
	return i, err
}

func myStrToInt2(s string) (int, error) {
	val := 0
	_, err := fmt.Sscanf(s, "%d", &val)
	return val, err
}