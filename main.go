package main

import (
	"fmt"
	convert "home5_1/convert"
)

func main() {
	s := "0,1254fg"
	i, err := convert.MyStrToInt(s)
	fmt.Println(i, err)
}

