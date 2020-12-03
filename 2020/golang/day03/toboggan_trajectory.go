package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	row := 0

	treeCounts := []int{0, 0, 0, 0, 0}
	for scanner.Scan() {
		// X position for each slope depends on the row
		// row 1, [1, 3, 5, 7, x]
		// row 2, [2, 6, 10, 14, 1]
		text := strings.TrimSpace(scanner.Text())
		textLength := len(text)

		pos1 := row
		pos2 := row * 3
		pos3 := row * 5
		pos4 := row * 7

		// last position increments every other
		pos5 := row / 2

		if string(text[pos1%textLength]) == "#" {
			treeCounts[0]++
		}

		if string(text[pos2%textLength]) == "#" {
			treeCounts[1]++
		}

		if string(text[pos3%textLength]) == "#" {
			treeCounts[2]++
		}

		if string(text[pos4%textLength]) == "#" {
			treeCounts[3]++
		}

		// Last position is only checked on even rows
		if (row%2) == 0 && string(text[pos5%textLength]) == "#" {
			treeCounts[4]++
		}

		row++
	}

	product := 1
	for idx, num := range treeCounts {
		product *= num
	}

	log.Printf("product is %d", product)
}
