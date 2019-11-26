package main

import (
	"fmt"
)

func main() {
	x1 := []int{1, 2}
	x2 := []int{10, 20}
	x3 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(average(x1))
	fmt.Println(average(x2))
	fmt.Println(average(x3))
}

func average(arr []int) float64 {
	sum := 0
	for _, v := range arr {
		sum += v
	}

	return float64(sum) / float64(len(arr))
}
