package attempt1

import (
	"fmt"
	"regexp"
	"strings"
)

var cache = make(map[string]string)

func puzzle1(input string) (result int) {
	split := strings.Split(input, "\n\n")
	messages := strings.Split(split[1], "\n")
	rules := getRules(strings.Split(split[0], "\n"))

	rule := simplifyRule(rules["0"], rules)

	longestRulePart := findLongestPart(rule)
	for _, message := range messages {
		if len(message) > len(longestRulePart) {
			messages = remove(messages, message)
		}
	}

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
		result = strings.Replace(result, match[0], fmt.Sprintf("(%s)", rules[index]), -1)
		rule = result
		matches = r.FindAllStringSubmatch(result, -1)
	}
	return expandBrackets(strings.Replace(result, " ", "", -1))
}

func expandBrackets(rule string) string {
	for strings.Contains(rule, ")(") {
		joinIndex := strings.Index(rule, ")(")
		openingBracketIndex := findOpeningBracketIndex(rule, joinIndex)
		closingBracketIndex := findClosingBracketIndex(rule, joinIndex+1)
		bracket1 := rule[openingBracketIndex+1 : joinIndex]
		bracket2 := rule[joinIndex+2 : closingBracketIndex]
		rule = rule[:openingBracketIndex+1] + expandPair(bracket1, bracket2) + rule[closingBracketIndex:]
	}
	return removeBrackets(rule)
}

func expandPair(bracket1 string, bracket2 string) string {
	if hit, ok := cache[bracket1+bracket2]; ok {
		return hit
	}

	if strings.Contains(bracket1, "(") {
		bracket1 = expandBrackets(bracket1)
	}
	if strings.Contains(bracket2, "(") {
		bracket2 = expandBrackets(bracket2)
	}

	flipped := false
	if strings.Contains(bracket1, "|") {
		flipped = true
		wrappedBracket2 := fmt.Sprintf("(%s)", bracket2)
		bracket2 = bracket1
		bracket1 = wrappedBracket2
	}

	split := strings.Split(bracket2, "|")

	for i, s := range split {
		if flipped {
			s = fmt.Sprintf("(%s)", s)
			split[i] = s + bracket1
		} else {
			split[i] = bracket1 + s
		}
	}
	result := strings.Join(split, "|")
	cache[bracket1+bracket2] = result
	return result
}

func findClosingBracketIndex(expression string, openingBracketIndex int) int {
	result := openingBracketIndex
	count := 1
	for count > 0 {
		result++
		if string(expression[result]) == "(" {
			count++
		} else if string(expression[result]) == ")" {
			count--
		}
	}
	return result
}

func findOpeningBracketIndex(expression string, closingBracketIndex int) int {
	result := closingBracketIndex
	count := 1
	for count > 0 {
		result--
		if string(expression[result]) == ")" {
			count++
		} else if string(expression[result]) == "(" {
			count--
		}
	}
	return result
}

func removeBrackets(rule string) string {
	return strings.Replace(strings.Replace(rule, "(", "", -1), ")", "", -1)
}

func findLongestPart(rule string) (result string) {
	for _, part := range strings.Split(rule, "|") {
		if len(part) > len(result) {
			result = part
		}
	}
	return
}

func remove(slice []string, x string) (result []string) {
	for _, v := range slice {
		if v != x {
			result = append(result, v)
		}
	}
	return
}

func checkMessage(message string, rule string) bool {
	r := regexp.MustCompile(rule)
	return r.MatchString(message)
}
