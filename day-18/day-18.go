package main

import (
	"fmt"
	"regexp"
	"strings"
	. "utils"
)

func puzzle1(input string) (result int) {
	return evaluateInput(input, evaluateLeftToRight)
}

func puzzle2(input string) int {
	return evaluateInput(input, evaluateAdditionBeforeMultiplication)
}

func evaluateInput(input string, evaluateExpression func(string) int) (result int) {
	expressions := strings.Split(input, "\n")
	for _, expression := range expressions {
		result += evaluateNestedExpression(expression, evaluateExpression)
	}
	return
}

func evaluateNestedExpression(expression string, evaluateExpression func(string) int) int {
	for strings.Contains(expression, "(") {
		r := regexp.MustCompile(`(\(([\d\*\+ ]*)\))`)
		matches := r.FindAllStringSubmatch(expression, -1)
		for _, match := range matches {
			subExpression := match[2]
			evaluatedSubExpression := evaluateNestedExpression(subExpression, evaluateExpression)
			expression = strings.Replace(expression, match[1], fmt.Sprint(evaluatedSubExpression), -1)
		}
	}
	return evaluateExpression(expression)
}

func evaluateLeftToRight(expression string) int {
	e := strings.Split(expression, " ")
	result := ConvertToInt(e[0])
	for i := 1; i < len(e)-1; i += 2 {
		if e[i] == "*" {
			result *= ConvertToInt(e[i+1])
			continue
		}
		result += ConvertToInt(e[i+1])
	}
	return result
}

func evaluateAdditionBeforeMultiplication(expression string) int {
	r := regexp.MustCompile(`(\d+ \+ \d+)`)
	matches := r.FindAllStringSubmatch(expression, -1)
	for len(matches) > 0 {
		match := matches[0]
		subExpression := match[1]
		expression = strings.Replace(expression, match[0], fmt.Sprint(evaluateLeftToRight(subExpression)), -1)
		matches = r.FindAllStringSubmatch(expression, -1)
	}
	return evaluateLeftToRight(expression)
}
