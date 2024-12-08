package day3

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func Day3Challenge1() {
	fmt.Println("Day3Challenge1")

	file, err := os.Open("./day3/challenge1.txt")
	if err != nil {
		fmt.Println("could not read file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
	mulRegex := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	for scanner.Scan() {
		line := scanner.Text()
		matches := mulRegex.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			digits, _ := StringsToIntegers(strings.Split(match[0][4:len(match[0])-1], ","))
			total += digits[0] * digits[1]
		}
	}
	fmt.Println(total)
}

func Day3Challenge2() {
	fmt.Println("Day3Challenge2")

	file, err := os.Open("./day3/challenge2.txt")
	if err != nil {
		fmt.Println("could not read file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
	mulRegex := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	dontRegex := regexp.MustCompile(`don't\(\)`)
	doRegex := regexp.MustCompile(`do\(\)`)
	enabled := true
	for scanner.Scan() {
		fmt.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$NEWLINE$$$$$$$$$$$$$$$$$$")
		line := scanner.Text()
		start := 0
		for start < len(line) {
			// fmt.Println("INFO", start, len(line), enabled)
			searchLine := line[start:]
			if enabled {
				mulIndex := mulRegex.FindStringSubmatchIndex(searchLine)
				dontIndex := dontRegex.FindStringSubmatchIndex(searchLine)
				if len(mulIndex) == 0 {
					break
				}
				if len(dontIndex) != 0 {
					if dontIndex[0] < mulIndex[0] {
						fmt.Println("SKIPPING")
						start += dontIndex[1]
						enabled = false
						continue
					}
				}
				start += mulIndex[1]
				digits, _ := StringsToIntegers(strings.Split(searchLine[mulIndex[0]+4:mulIndex[1]-1], ","))
				total += digits[0] * digits[1]
				fmt.Println("DIGITS", digits, "TOTAL:", total)
			} else {
				doIndex := doRegex.FindStringSubmatchIndex(searchLine)
				if len(doIndex) == 0 {
					break
				}
				start += doIndex[1]
				enabled = true
				fmt.Println("ENABLE")
			}
		}
	}
	fmt.Println("TOTAL:", total)
}
