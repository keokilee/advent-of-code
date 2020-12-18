package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	Start int
	End   int
}

func (r *Range) Cover(value int) bool {
	returnVal := value >= r.Start && value <= r.End
	return returnVal
}

func NewRange(rangeStr string) *Range {
	ranges := strings.Split(strings.TrimSpace(rangeStr), "-")
	start, _ := strconv.Atoi(ranges[0])
	end, _ := strconv.Atoi(ranges[1])
	return &Range{Start: start, End: end}
}

type Rule struct {
	Name   string
	Ranges []*Range
}

func NewRule(ruleStr string) *Rule {
	components := strings.Split(ruleStr, ":")
	name := components[0]
	rangeStrings := strings.Split(components[1], " or ")
	ranges := make([]*Range, 0)
	for _, rangeString := range rangeStrings {
		ranges = append(ranges, NewRange(rangeString))
	}

	return &Rule{
		Name:   name,
		Ranges: ranges,
	}
}

func (rule *Rule) Valid(value int) bool {
	for _, r := range rule.Ranges {
		if r.Cover(value) {
			return true
		}
	}

	return false
}

func NewTicket(ticketStr string) []int {
	nums := make([]int, 0)
	for _, numStr := range strings.Split(ticketStr, ",") {
		num, _ := strconv.Atoi(numStr)
		nums = append(nums, num)
	}

	return nums
}

func TicketErrors(ticket []int, rules []*Rule) int {
	errCode := 0
	for _, value := range ticket {
		valid := false
		for _, rule := range rules {
			if rule.Valid(value) {
				valid = true
				break
			}
		}

		if !valid {
			errCode += value
		}
	}

	return errCode
}

func main() {
	f, _ := os.Open("nearby_tickets.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	rules := makeRules()
	validTickets := make([][]int, 0)
	errCode := 0
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		ticket := NewTicket(text)
		errs := TicketErrors(ticket, rules)
		if errs == 0 {
			validTickets = append(validTickets, ticket)
		}

		errCode += errs
	}

	log.Printf("errcode=%d", errCode)
	log.Printf("valid tickets=%d", len(validTickets))
	ruleOrder(validTickets, rules)
}

func makeRules() []*Rule {
	f, _ := os.Open("rules.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	rules := make([]*Rule, 0)
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		rules = append(rules, NewRule(text))
	}

	return rules
}

func ruleOrder(validTickets [][]int, rules []*Rule) {
	// Determine rules that are valid for each column
	rulesForColumns := make([][]*Rule, 0)
	numRules := len(rules)
	for i := 0; i < numRules; i++ {
		column := ticketColumn(validTickets, i)

		validRules := make([]*Rule, 0)
		for _, rule := range rules {
			valid := true
			for _, value := range column {
				if !rule.Valid(value) {
					valid = false
					break
				}
			}

			if valid {
				validRules = append(validRules, rule)
			}
		}

		rulesForColumns = append(rulesForColumns, validRules)
	}

	// Find a permutation of rules that is valid
	for idx, column := range rulesForColumns {
		rules := []string{}
		for _, rule := range column {
			rules = append(rules, rule.Name)
		}

		log.Printf("col=%d rules=%s", idx, strings.Join(rules, ","))
	}

	rules = rulePermutation(rulesForColumns)
	for idx, rule := range rules {
		log.Printf("col %d rule %s", idx, rule.Name)
	}
}

func rulePermutation(columnRules [][]*Rule) []*Rule {
	selectedRules := make(map[*Rule]bool)
	rules, _ := findWithSelectedRules(columnRules, selectedRules)
	return rules
}

func findWithSelectedRules(columnRules [][]*Rule, selectedRules map[*Rule]bool) ([]*Rule, bool) {
	if len(columnRules) == 0 {
		return make([]*Rule, 0), true
	}

	ruleCol := columnRules[0]
	for _, rule := range ruleCol {
		if !selectedRules[rule] {
			selectedRules[rule] = true
			if rules, ok := findWithSelectedRules(columnRules[1:], selectedRules); ok {
				return append([]*Rule{rule}, rules...), true
			}
			selectedRules[rule] = false
		}
	}

	return nil, false
}

func ticketColumn(validTickets [][]int, index int) []int {
	nums := make([]int, 0)
	for _, ticket := range validTickets {
		nums = append(nums, ticket[index])
	}

	return nums
}
