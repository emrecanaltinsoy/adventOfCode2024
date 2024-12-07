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
	for scanner.Scan() {
		line := scanner.Text()
		r := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
		matches := r.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			digits, _ := StringsToIntegers(strings.Split(match[0][4:len(match[0])-1], ","))
			total += digits[0] * digits[1]
		}
	}
	fmt.Println(total)
}

func Day3Challenge2() {
	fmt.Println("Day3Challenge2")

	file, err := os.Open("./day2/challenge2.txt")
	if err != nil {
		fmt.Println("could not read file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		_, err := StringsToIntegers(strings.Split(line, " "))
		if err != nil {
			fmt.Println("Error converting to integer")
		}
	}
}
