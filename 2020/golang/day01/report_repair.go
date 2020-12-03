package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	keys := make(map[int]bool)
	nums := make([]int, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		num, _ := strconv.Atoi(text)
		keys[num] = true
		nums = append(nums, num)
	}

	for i, num := range nums {
		for _, num2 := range nums[i:] {
			remainder := 2020 - num - num2
			if _, ok := keys[remainder]; ok {
				fmt.Printf("Result is %d", num*num2*remainder)
				os.Exit(0)
			}
		}
	}

	fmt.Printf("no result found!")
	os.Exit(1)
}
