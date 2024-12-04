package day1

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"
)

func Day1Challenge1() {
	fmt.Println("Day1Challenge1")

	file, err := os.Open("./day1/challenge1.txt")
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

func Day1Challenge2() {
	fmt.Println("Day1Challenge2")

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
