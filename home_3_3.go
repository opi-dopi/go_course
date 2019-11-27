package main

import (
	"fmt"
	"sort"
)

func main() {
	a := map[int]string{10: "asd", 0: "qweqw", 500: "werwer"}
	b := map[int]string{2: "a", 0: "b", 1: "c"}

	printSorted(a)
	fmt.Println("-------------------------")
	printSorted(b)
}

func printSorted(x map[int]string) {

	arKeys := make([]int, 0, len(x))
	for k := range x {
		arKeys = append(arKeys, k)
	}

	sort.Ints(arKeys)

	for _, v := range arKeys {

		fmt.Printf("%v: ", v)
		fmt.Println(x[v])
	}

}
