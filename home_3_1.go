package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(average(slice))
}

func average(slice []int) float64 {
	summ := 0
	for _, val := range slice {
		summ += val
	}
	return float64(summ) / float64(len(slice))
}
