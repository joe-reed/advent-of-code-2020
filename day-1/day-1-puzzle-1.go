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
	lines := strings.Split(string(input), "\n")
	lines = lines[:len(lines)-1]

	for _, line1 := range lines {
		for _, line2 := range lines {
			int1 := ConvertToInt(line1)
			int2 := ConvertToInt(line2)

			sum := sum(int1, int2)
			product := product(int1, int2)
			if sum == 2020 {
				fmt.Println(product)
				return
			}
		}
	}
}

func sum(numbers ...int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}

func product(numbers ...int) int {
	result := 1
	for _, number := range numbers {
		result *= number
	}
	return result
}
