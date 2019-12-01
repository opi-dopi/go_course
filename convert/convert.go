package convert

import (
	"regexp"
	"strconv"
)

func MyStrToInt(s string) (int, error) {
	var re = regexp.MustCompile(`(?m)^\d+`)

	match := re.FindAllString(s, -1)
	if len(match) > 0 {
		i2, err := strconv.Atoi(match[0])
		return i2, err
	}
	return int(0), nil

}
