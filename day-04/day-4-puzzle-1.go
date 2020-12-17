package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	. "utils"
)

func main() {
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
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, requiredField := range requiredFields {
		if !strings.Contains(passport, requiredField) {
			return false
		}
	}
	return true
}
