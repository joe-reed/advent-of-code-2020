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

	for _, message := range messages {
		message, isValid := checkMessage(message, rules["0"], rules)
		if isValid && len(message) == 0 {
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

func checkMessage(message string, rule string, rules map[string]string) (string, bool) {
	isValid := false
	options := strings.Split(rule, "|")
	originalMessage := message
	for _, option := range options {
		optionValid := true
		for _, r := range strings.Split(option, " ") {
			rValid := true
			re := regexp.MustCompile(`(\d+)`)
			matches := re.FindAllStringSubmatch(r, -1)
			if len(matches) > 0 {
				match := matches[0]
				index := match[1]
				newMessage, valid := checkMessage(message, rules[index], rules)
				rValid = rValid && valid
				message = newMessage
			} else {
				rr := regexp.MustCompile(fmt.Sprintf("^%s", r))
				valid := rr.MatchString(message)
				rValid = rValid && valid
				if valid {
					message = strings.Replace(message, r, "", 1)
				}
			}
			optionValid = optionValid && rValid
		}
		isValid = isValid || optionValid
		if optionValid {
			break
		}
		message = originalMessage
	}

	return message, isValid
}
