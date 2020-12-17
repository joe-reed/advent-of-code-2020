package main

import (
	"fmt"
	"sort"
	"time"
	. "utils"
)

func main() {
	defer PrintTimeSince(time.Now())
	data := LoadFile("./input.txt")
	integers := []int{}
	for _, d := range data {
		integers = append(integers, ConvertToInt(d))
	}
	invalidNumber := getInvalidNumber(integers)
	r := findRangeSummingTo(invalidNumber, integers)
	sort.Ints(r)
	fmt.Println(r[0] + r[len(r)-1])
}

func getInvalidNumber(data []int) int {
	for i := 25; i < len(data); i++ {
		if !isSumOfTwoOfPrevious25(i, data) {
			return data[i]
		}
	}
	return 0
}

func isSumOfTwoOfPrevious25(i int, data []int) bool {
	value := data[i]

	for j := i - 25; j < i; j++ {
		for k := i - 25; k < i; k++ {
			if j == k {
				continue
			}
			if value == data[j]+data[k] {
				return true
			}
		}
	}
	return false
}

func findRangeSummingTo(total int, data []int) []int {
	length := 1
	r := []int{}
	for true {
		for i := 1; i < length; i++ {
			for j := 0; j < len(data); j++ {
				r = []int{}
				for k := 0; k < i; k++ {
					r = append(r, data[j+k])
				}
				if total == sum(r) {
					return r
				}
				if sum(r) > total {
					break
				}
			}
		}
		length++
	}
	return r
}

func sum(numbers []int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}
