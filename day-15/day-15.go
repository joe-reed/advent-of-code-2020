package main

import (
	"strings"
	. "utils"
)

func puzzle1(input string) int {
	return getNthSpoken(input, 2020)
}

func puzzle2(input string) int {
	return getNthSpoken(input, 30000000)
}

func getNthSpoken(input string, n int) int {
	startingNumbers := MapToInts(strings.Split(input, ","))
	appearances := make(map[int][]int)

	for i, startingNumber := range startingNumbers {
		appearances[startingNumber] = append(appearances[startingNumber], i)
	}

	lastSpoken := startingNumbers[len(startingNumbers)-1]
	for i := len(startingNumbers); i < n; i++ {
		spoken := 0

		appearanceCount := len(appearances[lastSpoken])
		if appearanceCount > 1 {
			spoken = appearances[lastSpoken][appearanceCount-1] - appearances[lastSpoken][appearanceCount-2]
		}

		appearances[spoken] = append(appearances[spoken], i)
		lastSpoken = spoken
	}

	return lastSpoken
}
