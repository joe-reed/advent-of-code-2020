package main

import (
	"bufio"
	"fmt"
	"os"
	. "utils"
)

func main() {
	file, err := os.Open("./input.txt")
	Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	currentHorizontalPosition := 0
	treeCount := 0
	lineLength := 0

	for scanner.Scan() {
		line := scanner.Text()
		lineLength = len(line)

		if currentHorizontalPosition > lineLength-1 {
			currentHorizontalPosition = currentHorizontalPosition - lineLength
		}

		if isLetterAtIndex(line, "#", currentHorizontalPosition) {
			treeCount++
		}

		currentHorizontalPosition += 3
	}

	fmt.Println(treeCount)
}

func isLetterAtIndex(password string, letter string, index int) bool {
	return string(password[index]) == letter
}
