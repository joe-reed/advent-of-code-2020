package main

import (
	"fmt"
	"strings"
	. "utils"
)

func main() {
	rules := LoadFile("./input.txt")
	options := getOptions(rules, "shiny gold")
	fmt.Println(len(options))
}

func getOptions(rules []string, bag string) []string {
	options := []string{}
	for _, rule := range rules {
		containingBag := strings.Split(rule, " bags")[0]

		if strings.Contains(rule, bag) && containingBag != bag {
			options = appendIfMissing(options, containingBag)

			for _, additionalOption := range getOptions(rules, containingBag) {
				options = appendIfMissing(options, additionalOption)
			}
		}
	}
	return options
}

func appendIfMissing(options []string, anOption string) []string {
	for _, option := range options {
		if option == anOption {
			return options
		}
	}
	return append(options, anOption)
}
