package day4

import (
	"strconv"
)

func StringsToIntegers(lines []string) ([]int, error) {
	integers := make([]int, 0, len(lines))
	for _, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		integers = append(integers, n)
	}
	return integers, nil
}

func AbsDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
