package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	. "utils"
)

func main() {
	input, err := ioutil.ReadFile("./input.txt")
	Check(err)
	answerGroups := strings.Split(string(input), "\n\n")

	totalCount := 0
	for _, answerGroup := range answerGroups {
		totalCount += getCount(answerGroup)
	}
	fmt.Println(totalCount)
}

func getCount(answerGroup string) int {
	answerSets := strings.Split(answerGroup, "\n")
	uniqueAnswers := ""
	for _, answerSet := range answerSets {
		for _, answer := range answerSet {
			if strings.Contains(uniqueAnswers, string(answer)) {
				continue
			}
			uniqueAnswers += string(answer)
		}
	}
	return len(uniqueAnswers)
}
