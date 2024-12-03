package day1

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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

func Day1Challenge1() {
	fmt.Println("Day1Challenge1")

	file, err := os.Open("./day1/input.txt")
	if err != nil {
		fmt.Println("could not read file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var firstList []int
	var secondList []int
	for scanner.Scan() {
		line := scanner.Text()
		locationIds, err := StringsToIntegers(strings.Split(line, "   "))
		if err != nil {
			fmt.Println("Error converting to integer")
		}
		firstList = append(firstList, locationIds[0])
		secondList = append(secondList, locationIds[1])
	}
	sort.Ints(firstList)
	sort.Ints(secondList)

	totalDistance := 0
	for i := range firstList {
		totalDistance += AbsDiffInt(firstList[i], secondList[i])
	}
	fmt.Println(totalDistance)
}
