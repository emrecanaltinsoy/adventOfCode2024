package day2

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func Day2Challenge1() {
	fmt.Println("Day2Challenge1")

	file, err := os.Open("./day2/challenge1.txt")
	if err != nil {
		fmt.Println("could not read file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		locationIds, err := StringsToIntegers(strings.Split(line, " "))
		if err != nil {
			fmt.Println("Error converting to integer")
		}
		var doesIncrease bool
		var isValid bool
		for index := range locationIds[:len(locationIds)-1] {
			difference := locationIds[index+1] - locationIds[index]
			if index == 0 {
				if difference == 0 || difference > 3 || difference < -3 {
					isValid = false
					break
				}
				if difference > 0 {
					doesIncrease = true
				} else {
					doesIncrease = false
				}
			} else {
				if difference == 0 || difference > 3 || difference < -3 {
					isValid = false
					break
				}
				if difference > 0 && !doesIncrease {
					isValid = false
					break
				} else if difference < 0 && doesIncrease {
					isValid = false
					break
				} else {
					isValid = true
				}
			}
		}
		if isValid {
			total++
		}
	}
	fmt.Println("total", total)
}

func Day2Challenge2() {
	fmt.Println("Day2Challenge2")

	file, err := os.Open("./day1/challenge2.txt")
	if err != nil {
		fmt.Println("could not read file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var firstList []int
	locationIDMap := make(map[int]int)
	for scanner.Scan() {
		line := scanner.Text()
		locationIds, err := StringsToIntegers(strings.Split(line, "   "))
		if err != nil {
			fmt.Println("Error converting to integer")
		}
		if !slices.Contains(firstList, locationIds[0]) {
			firstList = append(firstList, locationIds[0])
		}

		val, ok := locationIDMap[locationIds[1]]
		if ok {
			locationIDMap[locationIds[1]] = val + 1
		} else {
			locationIDMap[locationIds[1]] = 1
		}
	}
	total := 0
	for _, firstListItem := range firstList {
		total += firstListItem * locationIDMap[firstListItem]
	}

	fmt.Println(total)
}
