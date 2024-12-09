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
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	getCharacter := func(r int, c int) string {
		return string(lines[r][c])
	}
	R := len(lines)
	C := len(lines[0])
	for r := range R {
		for c := range C {
			if getCharacter(r, c) != "X" {
				continue
			}
			fmt.Println(getCharacter(r, c))
			// Left-Right
			if c+3 < C && getCharacter(r, c+1) == "M" && getCharacter(r, c+2) == "A" && getCharacter(r, c+3) == "S" {
				fmt.Println("Left-Right")
				total++
			}
			// Right-Left
			if c-3 >= 0 && getCharacter(r, c-1) == "M" && getCharacter(r, c-2) == "A" && getCharacter(r, c-3) == "S" {
				fmt.Println("Right-Left")
				total++
			}
			// Top-Bottom
			if r+3 < R && getCharacter(r+1, c) == "M" && getCharacter(r+2, c) == "A" && getCharacter(r+3, c) == "S" {
				fmt.Println("Top-Bottom")
				total++
			}
			// Bottom-Top
			if r-3 >= 0 && getCharacter(r-1, c) == "M" && getCharacter(r-2, c) == "A" && getCharacter(r-3, c) == "S" {
				fmt.Println("Bottom-Top")
				total++
			}
			// Diagonal lefttop-rightbottom
			if c+3 < C && r+3 < R && getCharacter(r+1, c+1) == "M" && getCharacter(r+2, c+2) == "A" && getCharacter(r+3, c+3) == "S" {
				fmt.Println("Diagonal lefttop-rightbottom")
				total++
			}
			// Diagonal leftbottom-righttop
			if c+3 < C && r-3 >= 0 && getCharacter(r-1, c+1) == "M" && getCharacter(r-2, c+2) == "A" && getCharacter(r-3, c+3) == "S" {
				fmt.Println("Diagonal leftbottom-righttop")
				total++
			}
			// Diagonal rightbottom-lefttop
			if c-3 >= 0 && r-3 >= 0 && getCharacter(r-1, c-1) == "M" && getCharacter(r-2, c-2) == "A" && getCharacter(r-3, c-3) == "S" {
				fmt.Println("Diagonal rightbottom-lefttop")
				total++
			}
			// Diagonal righttop-leftbottom
			if c-3 >= 0 && r+3 < R && getCharacter(r+1, c-1) == "M" && getCharacter(r+2, c-2) == "A" && getCharacter(r+3, c-3) == "S" {
				fmt.Println("Diagonal righttop-leftbottom")
				total++
			}
		}
	}
	fmt.Printf("TOTAL: %d\n", total)
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
