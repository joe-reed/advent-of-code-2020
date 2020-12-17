package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"sync"
	"time"
	. "utils"
)

func main() {
	defer PrintTimeSince(time.Now())

	input, err := ioutil.ReadFile("./input.txt")
	Check(err)
	passports := strings.Split(string(input), "\n\n")

	c := make(chan bool, len(passports))
	var wg sync.WaitGroup

	for _, passport := range passports {
		wg.Add(1)
		go isValidPassport(passport, c, &wg)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	validPassportCount := 0
	for isValid := range c {
		if isValid {
			validPassportCount++
		}
	}
	fmt.Println(validPassportCount)
}

func isValidPassport(passport string, c chan<- bool, wg *sync.WaitGroup) {
	defer wg.Done()
	fieldRules := getFieldRules()
	requiredFields := GetKeys(fieldRules)
	for _, requiredField := range requiredFields {
		if !strings.Contains(passport, requiredField) {
			c <- false
			return
		}
	}

	for field, rule := range fieldRules {
		value := extractValueForField(passport, field)
		if !rule(value) {
			c <- false
			return
		}
	}

	c <- true
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
