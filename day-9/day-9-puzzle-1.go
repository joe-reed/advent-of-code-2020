package main

import (
	"fmt"
	. "utils"
)

func main() {
	data := LoadFile("./input.txt")
	for i := 25; i < len(data); i++ {
		if !isSumOfTwoOfPrevious25(i, data) {
			fmt.Println(data[i])
			break
		}
	}
}

func isSumOfTwoOfPrevious25(i int, data []string) bool {
	value := ConvertToInt(data[i])

	for j := i - 25; j < i; j++ {
		for k := i - 25; k < i; k++ {
			if j == k {
				continue
			}
			if value == ConvertToInt(data[j])+ConvertToInt(data[k]) {
				return true
			}
		}
	}
	return false
}
