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
	answerGroups := strings.Split(string(input), "\n\n")

	totalCount := 0
	for _, answerGroup := range answerGroups {
		totalCount += getCount(answerGroup)
	}
	fmt.Println(totalCount)
}

func getCount(answerGroup string) int {
	answerSets := strings.Split(answerGroup, "\n")
	answeredByAll := ""
	firstAnswerSet := answerSets[0]
	for _, answer := range firstAnswerSet {
		answered := true

		for _, answerSet := range answerSets {
			if answerSet == "" {
				continue
			}
			if !strings.Contains(answerSet, string(answer)) {
				answered = false
				break
			}
		}

		if answered {
			answeredByAll += string(answer)
		}
	}

	return len(answeredByAll)
}
