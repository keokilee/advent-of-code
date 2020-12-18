package main

import (
	"log"
	"strconv"
	"strings"
)

const input = "2,0,6,12,1,3"
const turns = 30000000

func main() {
	startingNumbers := getStartingNumbers()

	spokenNumbers := make([]int, 0)
	indexes := make(map[int][]int)
	for i := 0; i < turns; i++ {
		if i < len(startingNumbers) {
			value := startingNumbers[i]
			spokenNumbers = append(spokenNumbers, value)
			indexes[value] = []int{i}
		} else {
			lastSpoken := spokenNumbers[i-1]

			previous, ok := previouslySpoken(indexes, lastSpoken)
			if !ok {
				// log.Printf("turn=%d previous=%d next=%d", i, lastSpoken, 0)
				spokenNumbers = append(spokenNumbers, 0)
			} else {
				spokenNumbers = append(spokenNumbers, i-1-previous)
				// log.Printf("turn=%d previous=%d next=%d", i, lastSpoken, i-1-previous)
			}

			// Update indexes
			if _, ok := indexes[lastSpoken]; ok {
				indexes[lastSpoken] = append(indexes[lastSpoken], i-1)
			} else {
				indexes[lastSpoken] = []int{i - 1}
			}
		}
	}

	lastIdx := len(spokenNumbers) - 1
	lastSpoken := spokenNumbers[lastIdx]
	log.Printf("last value=%d", lastSpoken)
}

func previouslySpoken(indexes map[int][]int, value int) (int, bool) {
	indices, ok := indexes[value]
	if !ok {
		return 0, false
	}

	last := len(indices) - 1
	return indices[last], true
}

func getStartingNumbers() []int {
	numStr := strings.Split(input, ",")

	nums := make([]int, 0)
	for _, str := range numStr {
		num, _ := strconv.Atoi(str)
		nums = append(nums, num)
	}

	return nums
}
