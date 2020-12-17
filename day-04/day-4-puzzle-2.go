package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
	. "utils"
)

func main() {
	defer PrintTimeSince(time.Now())

	input, err := ioutil.ReadFile("./input.txt")
	Check(err)
	passports := strings.Split(string(input), "\n\n")

	validPassportCount := 0
	for _, passport := range passports {
		if isValidPassport(passport) {
			validPassportCount++
		}
	}
	fmt.Println(validPassportCount)
}

func isValidPassport(passport string) bool {
	fieldRules := getFieldRules()
	requiredFields := GetKeys(fieldRules)
	for _, requiredField := range requiredFields {
		if !strings.Contains(passport, requiredField) {
			return false
		}
	}

	for field, rule := range fieldRules {
		value := extractValueForField(passport, field)
		if !rule(value) {
			return false
		}
	}

	return true
}

func getFieldRules() map[string]func(string) bool {
	fieldRules := make(map[string]func(string) bool)
	fieldRules["byr"] = func(value string) bool {
		integerValue := ConvertToInt(value)
		return 1920 <= integerValue && integerValue <= 2002
	}

	fieldRules["iyr"] = func(value string) bool {
		integerValue := ConvertToInt(value)
		return 2010 <= integerValue && integerValue <= 2020
	}

	fieldRules["eyr"] = func(value string) bool {
		integerValue := ConvertToInt(value)
		return 2020 <= integerValue && integerValue <= 2030
	}

	fieldRules["hgt"] = func(value string) bool {
		unit := value[len(value)-2:]
		amount := value[:len(value)-2]
		if unit != "cm" && unit != "in" {
			return false
		}

		integerAmount := ConvertToInt(amount)
		if unit == "cm" {
			return 150 <= integerAmount && integerAmount <= 193
		}
		return 59 <= integerAmount && integerAmount <= 76
	}

	fieldRules["hcl"] = func(value string) bool {
		return CheckRegexp("^#[0-9a-f]{6}$", value)
	}

	fieldRules["ecl"] = func(value string) bool {
		validEyeColours := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
		return StringInArray(value, validEyeColours)
	}

	fieldRules["pid"] = func(value string) bool {
		return CheckRegexp("^[0-9]{9}$", value)
	}

	return fieldRules
}

func extractValueForField(passport string, field string) string {
	split := strings.Split(passport, field+":")
	portionStartingWithValue := split[1]
	splitAgain := strings.Split(portionStartingWithValue, "\n")
	value := splitAgain[0]

	if strings.Contains(value, " ") {
		splitAgain = strings.Split(portionStartingWithValue, " ")
		value = splitAgain[0]
	}

	return value
}
