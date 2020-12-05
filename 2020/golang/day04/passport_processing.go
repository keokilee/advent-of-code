package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	passportBuf := []string{}
	validCount := 0
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())

		// Check for password delimiter
		if text == "" {
			if validPassport(passportBuf) {
				validCount++
			}
			passportBuf = []string{}
		} else {
			passportBuf = append(passportBuf, text)
		}
	}

	// Need to validate last password at EOF
	if len(passportBuf) > 0 && validPassport(passportBuf) {
		validCount++
	}

	log.Printf("Valid passport count: %d", validCount)
}

var validators = map[string]func(string) bool{
	"byr": func(year string) bool {
		if len(year) != 4 {
			return false
		}

		num, err := strconv.Atoi(year)
		if err != nil {
			return false
		}

		return num >= 1920 && num <= 2002
	},
	"iyr": func(year string) bool {
		if len(year) != 4 {
			return false
		}

		num, err := strconv.Atoi(year)
		if err != nil {
			return false
		}

		return num >= 2010 && num <= 2020
	},
	"eyr": func(year string) bool {
		if len(year) != 4 {
			return false
		}

		num, err := strconv.Atoi(year)
		if err != nil {
			return false
		}

		return num >= 2020 && num <= 2030
	},
	"hgt": func(height string) bool {
		if !strings.HasSuffix(height, "in") && !strings.HasSuffix(height, "cm") {
			return false
		}

		if strings.HasSuffix(height, "in") {
			value := strings.TrimSuffix(height, "in")
			num, err := strconv.Atoi(value)
			if err != nil {
				return false
			}

			return num >= 59 && num <= 76
		}

		value := strings.TrimSuffix(height, "cm")
		num, err := strconv.Atoi(value)
		if err != nil {
			return false
		}

		return num >= 150 && num <= 193
	},
	"hcl": func(color string) bool {
		matched, err := regexp.MatchString("^#[0-9a-f]{6}$", color)
		if err != nil {
			panic(err)
		}

		return matched
	},
	"ecl": func(color string) bool {
		validColors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
		for _, validColor := range validColors {
			if validColor == color {
				return true
			}
		}

		return false
	},
	"pid": func(id string) bool {
		matched, err := regexp.MatchString("^[0-9]{9}$", id)
		if err != nil {
			panic(err)
		}

		return matched
	},
	"cid": func(string) bool {
		return true
	},
}

func validPassport(lines []string) bool {
	validations := map[string]bool{"byr": false, "iyr": false, "eyr": false, "hgt": false, "hcl": false, "ecl": false, "pid": false, "cid": true}
	passport := strings.Join(lines, " ")
	components := strings.Split(passport, " ")

	for _, field := range components {
		fieldComponents := strings.Split(field, ":")
		field, value := fieldComponents[0], fieldComponents[1]
		validations[field] = validators[field](value)
	}

	for _, valid := range validations {
		if !valid {
			return false
		}
	}

	return true
}
