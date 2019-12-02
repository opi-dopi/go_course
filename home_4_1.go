package main

import (
	"fmt"
	"os"
	"sort"
	"time"
)

type Person struct {
	firstName string
	lastName  string
	birthDay  time.Time
}

type People []Person

func (p People) Len() int {
	return len(p)
}

func (p People) Less(i, j int) bool {

	if p[j].birthDay.Before(p[i].birthDay) {
		return true
	}

	if p[j].birthDay.Equal(p[i].birthDay) && p[j].firstName > p[i].firstName {
		return true
	}

	return false
}

func (p People) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s ( %s )", p.firstName, p.lastName, p.birthDay.Format("2006-01-02"))
}

func main() {
	const layout= "2006-01-02"
	ivanIvanovDate1, err := time.Parse(layout, "2005-08-10")
	ivanIvanovDate2, err := time.Parse(layout, "2003-08-05")
	ivanIvanovDate3, err := time.Parse(layout, "2005-08-10")
	ivanIvanovDate4, err := time.Parse(layout, "2001-08-10")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	p := People{
		{"Ivan", "Ivanov", ivanIvanovDate1},
		{"Ivan", "Ivanov", ivanIvanovDate2},
		{"Artiom", "Ivanov", ivanIvanovDate3},
		{"Artiom", "Pechka", ivanIvanovDate4},
	}

	sort.Sort(p)

	fmt.Println(p)

	for _, v := range p {
		fmt.Println(v.String())
	}
}
