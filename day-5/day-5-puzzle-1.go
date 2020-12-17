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

	max := 0
	for scanner.Scan() {
		id := calculateId(scanner.Text())
		if id > max {
			max = id
		}
	}
	fmt.Println(max)
}

func calculateId(boardingPass string) int {
	row := calculateRow(boardingPass[:7])
	column := calculateColumn(boardingPass[7:])
	return row*8 + column
}

func calculateRow(rowSpecifier string) int {
	return evaluateLocation(rowSpecifier, 128, "F")
}

func calculateColumn(columnSpecifier string) int {
	return evaluateLocation(columnSpecifier, 8, "L")
}

func evaluateLocation(specifier string, spaceSize int, startHalfCharacter string) int {
	start, end := 0, spaceSize
	for _, char := range specifier {
		average := (end + start) / 2
		if string(char) == startHalfCharacter {
			end = average
		} else {
			start = average
		}
	}
	return end - 1
}
