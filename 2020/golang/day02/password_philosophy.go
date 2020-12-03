package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type passwordPolicy struct {
	LowerBound int
	UpperBound int
	Letter     string
}

func newPasswordPolicy(policyString string) *passwordPolicy {
	components := strings.Split(policyString, " ")
	rangeComponents := strings.Split(components[0], "-")
	lowerBound, _ := strconv.Atoi(rangeComponents[0])
	upperBound, _ := strconv.Atoi(rangeComponents[1])

	return &passwordPolicy{
		LowerBound: lowerBound,
		UpperBound: upperBound,
		Letter:     components[1],
	}
}

func (policy *passwordPolicy) IsValidPassword(password string) bool {
	stripped := strings.Replace(password, policy.Letter, "", -1)
	diff := len(password) - len(stripped)

	return diff >= policy.LowerBound && diff <= policy.UpperBound
}

func (policy *passwordPolicy) IsValidPassword2(password string) bool {
	pos1Valid := string(password[policy.LowerBound-1]) == policy.Letter
	pos2Valid := string(password[policy.UpperBound-1]) == policy.Letter

	return pos1Valid != pos2Valid
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	validPasswordCount := 0
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		components := strings.Split(text, ":")
		policy := newPasswordPolicy(components[0])
		password := strings.TrimSpace(components[1])
		if policy.IsValidPassword2(password) {
			validPasswordCount++
		}
	}

	fmt.Printf("Valid password count: %d", validPasswordCount)
}
