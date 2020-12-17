package main

import (
	"fmt"
	"regexp"
	"strings"
	. "utils"
)

func main() {
	rules := LoadFile("./input.txt")
	fmt.Println(getSizeOfContents(rules, "shiny gold") - 1)
}

func getSizeOfContents(rules []string, bag string) int {
	sizeOfContents := 1
	re := regexp.MustCompile(`\d+`)
	for _, rule := range rules {
		containingBag := strings.Split(rule, " bags")[0]
		if containingBag != bag {
			continue
		}

		numbers := re.FindAll([]byte(rule), -1)
		if numbers == nil {
			break
		}

		for _, number := range numbers {
			innerBagRe := regexp.MustCompile(string(number) + " (.*?) bag")
			innerBag := innerBagRe.FindStringSubmatch(rule)[1]

			rule = strings.Replace(rule, string(number), "", 1)
			sizeOfContents += ConvertToInt(string(number)) * getSizeOfContents(rules, innerBag)
		}
	}
	return sizeOfContents
}
