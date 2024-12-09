package day5

import (
	"bufio"
	"fmt"
	"os"
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
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	fmt.Printf("total: %d\n", total)
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
