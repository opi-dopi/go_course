package convert

import (
	"testing"
)

type testpair struct {
	str    string
	number int
}

var tests = []testpair{
	{"12qwe", 12},
	{"014", 14},
	{"weqwer23", 0},
	{"1,4", 1},
}

func TestMyStrToInt(t *testing.T) {
	for _, pair := range tests {
		v, _ := MyStrToInt(pair.str)
		if v != pair.number {
			t.Error(
				"For", pair.str,
				"expected", pair.number,
				"got", v,
			)
		}
	}
}
