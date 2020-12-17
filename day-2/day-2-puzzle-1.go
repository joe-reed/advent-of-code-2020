package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	. "utils"
)

func main() {
	file, err := os.Open("./input.txt")
	Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	validPasswordCount := 0
	for scanner.Scan() {
		rule, password := parseLine(scanner.Text())
		isValid := checkPasswordAgainstRule(password, rule)
		if isValid {
			validPasswordCount++
		}
	}
	fmt.Println(validPasswordCount)
}

func parseLine(line string) (string, string) {
	arr := strings.Split(line, ": ")
	return arr[0], arr[1]
}

func checkPasswordAgainstRule(password string, rule string) bool {
	letter, min, max := parseRule(rule)

	letterCount := strings.Count(password, letter)
	return min <= letterCount && letterCount <= max
}

func parseRule(rule string) (string, int, int) {
	arr := strings.Split(rule, " ")
	letterCountRange := arr[0]
	letter := arr[1]

	arr = strings.Split(letterCountRange, "-")
	min := ConvertToInt(arr[0])
	max := ConvertToInt(arr[1])
	return letter, min, max
}
