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
	letter, position1, position2 := parseRule(rule)

	position1Valid := isLetterAtIndex(password, letter, position1-1)
	position2Valid := isLetterAtIndex(password, letter, position2-1)
	return (position1Valid || position2Valid) && !(position1Valid && position2Valid)
}

func isLetterAtIndex(password string, letter string, index int) bool {
	return string(password[index]) == letter
}

func parseRule(rule string) (string, int, int) {
	arr := strings.Split(rule, " ")
	letterCountRange := arr[0]
	letter := arr[1]

	arr = strings.Split(letterCountRange, "-")
	position1 := ConvertToInt(arr[0])
	position2 := ConvertToInt(arr[1])
	return letter, position1, position2
}
