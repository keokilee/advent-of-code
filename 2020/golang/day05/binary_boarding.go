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

	seats := make([]bool, 977)
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		row := getRow(text[:7], 0, 127)
		column := getColumn(text[7:], 0, 7)
		seatID := (row * 8) + column
		seats[seatID] = true
	}

	for id, occupied := range seats {
		if !occupied {
			log.Printf("available seat %d", id)
		}
	}
}

func getRow(text string, start, end int) int {
	if len(text) == 1 {
		if text == "F" {
			return start
		}

		return end
	}

	middle := start + ((end - start) / 2)
	char := string(text[0])

	if char == "F" {
		return getRow(text[1:], start, middle)
	}

	return getRow(text[1:], middle+1, end)
}

func getColumn(text string, start, end int) int {
	if len(text) == 1 {
		if text == "L" {
			return start
		}

		return end
	}

	middle := start + ((end - start) / 2)
	char := string(text[0])

	if char == "L" {
		return getColumn(text[1:], start, middle)
	}

	return getColumn(text[1:], middle+1, end)
}
