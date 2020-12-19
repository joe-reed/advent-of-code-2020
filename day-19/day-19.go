package main

import (
	"fmt"
	"regexp"
	"strings"
)

func puzzle1(input string) (result int) {
	split := strings.Split(input, "\n\n")
	messages := strings.Split(split[1], "\n")
	rules := getRules(strings.Split(split[0], "\n"))

	rule := simplifyRule(rules["0"], rules)

	for _, message := range messages {
		if checkMessage(message, rule) {
			result++
		}
	}
	return
}

func getRules(rules []string) map[string]string {
	result := make(map[string]string)
	for _, rule := range rules {
		split := strings.Split(rule, ": ")
		result[split[0]] = strings.Trim(split[1], "\" ")
	}
	return result
}

func simplifyRule(rule string, rules map[string]string) string {
	r := regexp.MustCompile(`(\d+)`)
	result := rule
	matches := r.FindAllStringSubmatch(result, -1)
	for len(matches) > 0 {
		match := matches[0]
		index := match[1]
		if r.MatchString(rules[index]) {
			rules[index] = simplifyRule(rules[index], rules)
		}
		result = strings.Replace(result, match[0], fmt.Sprintf("(%s)", rules[index]), 1)
		matches = r.FindAllStringSubmatch(result, -1)
	}
	return strings.Replace(result, " ", "", -1)
}

func checkMessage(message string, rule string) bool {
	r := regexp.MustCompile(fmt.Sprintf("^%s$", rule))
	return r.MatchString(message)
}
