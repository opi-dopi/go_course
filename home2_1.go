package main

import "fmt"

func main() {
	fmt.Println(factorial(6))
}

func factorial(i uint) uint {
	var fact uint = 1
	for i > 0 {
		fact *= i
		i--
	}
	return fact
}
