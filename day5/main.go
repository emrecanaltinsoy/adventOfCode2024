package day5

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func Day5() {
	fmt.Println("Day5")

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

	secondTotal := 0
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
		} else {
			// Create a map with initial indices and numbers
			pageIndexMap := make(map[int]int)
			for index := range len(pageNums) - 1 {
				for secondInd := index + 1; secondInd < len(pageNums); secondInd++ {
					_, ok := pageIndexMap[pageNums[index]]
					if !ok {
						pageIndexMap[pageNums[index]] = index
					}
					_, okSecond := pageIndexMap[pageNums[secondInd]]
					if !okSecond {
						pageIndexMap[pageNums[secondInd]] = secondInd
					}
					if !slices.Contains(pagesMap[pageNums[index]], pageNums[secondInd]) {
						// If next num is not in list of the first one, increase and decrease indices accordingly
						pageIndexMap[pageNums[index]]++
						pageIndexMap[pageNums[secondInd]]--
					}
				}
			}
			for num, ind := range pageIndexMap {
				if ind == len(pageNums)/2 {
					secondTotal += num
					break
				}
			}
		}
	}
	fmt.Printf("TOTAL: %d\n", total)
	fmt.Printf("SECOND TOTAL: %d\n", secondTotal)
}
