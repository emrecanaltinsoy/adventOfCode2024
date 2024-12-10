package day5

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func Day5Challenge1() {
	fmt.Println("Day5Challenge1")

	file, err := os.Open("./day5/input.txt")
	if err != nil {
		fmt.Println("could not read file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
	pagesMap := make(map[int][]int)
	var pageOrders []string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {
			pageNums, err := StringsToIntegers(strings.Split(line, "|"))
			if err != nil {
				fmt.Println("Error converting to integer")
			}
			_, ok := pagesMap[pageNums[0]]
			if ok {
				mappedPages := pagesMap[pageNums[0]]
				mappedPages = append(mappedPages, pageNums[1])
				pagesMap[pageNums[0]] = mappedPages
			} else {
				var mappedPages []int
				mappedPages = append(mappedPages, pageNums[1])
				pagesMap[pageNums[0]] = mappedPages
			}
		}
		if strings.Contains(line, ",") {
			pageOrders = append(pageOrders, line)
		}
	}
	for _, pageOrder := range pageOrders {
		pageNums, _ := StringsToIntegers(strings.Split(pageOrder, ","))
		isValid := true
		for index := range len(pageNums) - 1 {
			if !slices.Contains(pagesMap[pageNums[index]], pageNums[index+1]) {
				isValid = false
				break
			}
		}
		if isValid {
			total += pageNums[len(pageNums)/2]
		}
	}
	fmt.Printf("TOTAL: %d\n", total)
}

func Day5Challenge2() {
	fmt.Println("Day5Challenge2")

	file, err := os.Open("./day5/input.txt")
	if err != nil {
		fmt.Println("could not read file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	fmt.Printf("total: %d\n", total)
}
