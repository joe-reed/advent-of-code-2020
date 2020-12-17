package main

import (
	"bufio"
	"fmt"
	"os"
	. "utils"
)

func main() {
	count1 := countTreesForSlope(1, 1)
	count2 := countTreesForSlope(3, 1)
	count3 := countTreesForSlope(5, 1)
	count4 := countTreesForSlope(7, 1)
	count5 := countTreesForSlope(1, 2)

	fmt.Println(count1 * count2 * count3 * count4 * count5)
}

func isLetterAtIndex(password string, letter string, index int) bool {
	return string(password[index]) == letter
}

func countTreesForSlope(x int, y int) int {
	file, err := os.Open("./input.txt")
	Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	currentHorizontalPosition := 0
	treeCount := 0
	currentLine := -1
	lineLength := 0

	for scanner.Scan() {
		currentLine += 1
		if currentLine%y != 0 {
			continue
		}

		line := scanner.Text()
		lineLength = len(line)

		if currentHorizontalPosition > lineLength-1 {
			currentHorizontalPosition = currentHorizontalPosition - lineLength
		}

		if isLetterAtIndex(line, "#", currentHorizontalPosition) {
			treeCount++
		}

		currentHorizontalPosition += x
	}

	return treeCount
}
