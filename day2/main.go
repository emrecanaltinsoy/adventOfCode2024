package day2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func IsBadLevel(difference int) bool {
	return difference == 0 || difference > 3 || difference < -3
}

func All[T any](ts []T, pred func(T) bool) bool {
	for _, t := range ts {
		if !pred(t) {
			return false
		}
	}
	return true
}

func Day2Challenge1() {
	fmt.Println("Day2Challenge1")

	file, err := os.Open("./day2/challenge1.txt")
	if err != nil {
		fmt.Println("could not read file")
	}
	defer file.Close()

	IsPositive := func(value int) bool { return value > 0 }
	IsNegative := func(value int) bool { return value < 0 }

	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		locationIds, err := StringsToIntegers(strings.Split(line, " "))
		if err != nil {
			fmt.Println("Error converting to integer")
		}
		isValid := true
		var differences []int
		for index := range locationIds[:len(locationIds)-1] {
			differences = append(differences, locationIds[index+1]-locationIds[index])
		}
		if !All(differences, IsPositive) && !All(differences, IsNegative) {
			isValid = false
			continue
		}

		for _, diff := range differences {
			if IsBadLevel(diff) {
				isValid = false
				break
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

	file, err := os.Open("./day2/challenge2.txt")
	if err != nil {
		fmt.Println("could not read file")
	}
	defer file.Close()

	IsPositive := func(value int) bool { return value > 0 }
	IsNegative := func(value int) bool { return value < 0 }

	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		locationIds, err := StringsToIntegers(strings.Split(line, " "))
		if err != nil {
			fmt.Println("Error converting to integer")
		}
		isValid := true
		var differences []int
		for index := range locationIds[:len(locationIds)-1] {
			differences = append(differences, locationIds[index+1]-locationIds[index])
		}
		if !All(differences, IsPositive) && !All(differences, IsNegative) {
			isValid = false
			continue
		}

		for _, diff := range differences {
			if IsBadLevel(diff) {
				isValid = false
				break
			}
		}
		if isValid {
			total++
		}
	}
	fmt.Println("total", total)
}
