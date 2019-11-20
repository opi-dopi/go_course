package main

import (
	"fmt"
)

func main() {
	slice1 := []int{1, 2, 3, 4, 5, 6}
	//slice2 := []int{1, 2, 3, 5, 6, 7, 7, 8, 9, 22, 33}
	fmt.Println(average(slice1))
}

func average(slice []int) float64 {
	var summ = 0
	for _, v := range slice {
		fmt.Println(v)
		summ += v
	}
	return float64(summ) / float64(len(slice))
}
