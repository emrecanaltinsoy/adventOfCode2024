package day4

import (
	"bufio"
	"fmt"
	"os"
)

func Day4Challenge1() {
	fmt.Println("Day4Challenge1")

	file, err := os.Open("./day4/challenge1.txt")
	if err != nil {
		fmt.Println("could not read file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	fmt.Println(total)
}

func Day4Challenge2() {
	fmt.Println("Day4Challenge2")

	file, err := os.Open("./day4/challenge2.txt")
	if err != nil {
		fmt.Println("could not read file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	fmt.Println(total)
}
