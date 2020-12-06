package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Group struct {
	AnsweredQuestions map[string]int
	NumPeople         int
}

func NewGroup() *Group {
	return &Group{AnsweredQuestions: make(map[string]int)}
}

func (group *Group) RecordAnswers(answers string) {
	group.NumPeople += 1

	for _, char := range answers {
		_, ok := group.AnsweredQuestions[string(char)]
		if ok {
			group.AnsweredQuestions[string(char)]++
		} else {
			group.AnsweredQuestions[string(char)] = 1
		}
	}
}

func (group *Group) TotalAnsweredQuestions() int {
	total := 0
	for _, count := range group.AnsweredQuestions {
		if count == group.NumPeople {
			total++
		}
	}

	return total
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	group := NewGroup()
	total := 0
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			total += group.TotalAnsweredQuestions()
			group = NewGroup()
		} else {
			group.RecordAnswers(text)
		}
	}

	// Handle last group
	total += group.TotalAnsweredQuestions()

	log.Printf("Total is %d", total)
}
