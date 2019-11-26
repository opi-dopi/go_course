package main

import (
	"errors"
	"fmt"
)

func main() {
	var x1 []int
	x2 := []int{10, 20}
	x3 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(average(x1))
	fmt.Println(average(x2))
	fmt.Println(average(x3))
}

func average(arr []int) (float64, error) {

	if len(arr) == 0 {
		return float64(0), errors.New("Array length cannot be 0")
	}

	sum := 0
	for _, v := range arr {
		sum += v
	}

	return float64(sum) / float64(len(arr)), nil
}
