package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	sliceWords := []string{"one", "thr22", "two", "three3333", "thr33"}
	fmt.Println(max(sliceWords))

	sliceNum := []int64{0, 1, 2, 4, 5, 67, 8}
	fmt.Println(reverse(sliceNum))

}

func max(slice []string) string {
	lenSrt := 0
	var elem string
	for _, v := range slice {

		lenSrt = utf8.RuneCountInString(v)

		if utf8.RuneCountInString(elem) < lenSrt {
			elem = v
		}
	}

	return elem
}

func reverse(slice []int64) []int64 {
	sliceNum := make([]int64, len(slice))
	for j, i := 0, len(slice)-1; i >= 0; j, i = j+1, i-1 {

		sliceNum[j] = slice[i]
	}

	return sliceNum
}
