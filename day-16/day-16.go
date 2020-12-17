package main

import (
	"strings"
	. "utils"
)

func puzzle1(input string) int {
	rulesString, _, nearbyTicketsString := parseInput(input)

	ranges := getRanges(rulesString)
	nearbyTickets := getNearbyTickets(nearbyTicketsString)
	invalid := getInvalidNearbyTickets(nearbyTickets, ranges)

	result := 0
	for _, v := range invalid {
		result += v
	}
	return result
}

func puzzle2(input string) int {
	rulesString, myTicketString, nearbyTicketsString := parseInput(input)

	fieldRanges := getRanges(rulesString)
	nearbyTickets := getNearbyTickets(nearbyTicketsString)
	invalid := getInvalidNearbyTickets(nearbyTickets, fieldRanges)
	validNearbyTickets := removeTicketsContainingValues(nearbyTickets, invalid)

	fieldPositionOptions := make(map[string][]int)

	for field, ranges := range fieldRanges {
		fieldPositionOptions[field] = []int{}
		for i := 0; i < len(validNearbyTickets[0]); i++ {
			isOption := true
			for _, ticket := range validNearbyTickets {
				v := ticket[i]
				if !isInRange(v, ranges[0]) && !isInRange(v, ranges[1]) {
					isOption = false
				}
			}
			if isOption && !contains(fieldPositionOptions[field], i) {
				fieldPositionOptions[field] = append(fieldPositionOptions[field], i)
			}
		}
	}

	fieldPositions := make(map[string]int)
	for len(fieldPositions) != len(fieldRanges) {
		for field := range fieldRanges {
			if len(fieldPositionOptions[field]) == 1 {
				position := fieldPositionOptions[field][0]
				fieldPositions[field] = position
				for f, options := range fieldPositionOptions {
					if contains(options, position) {
						fieldPositionOptions[f] = remove(options, position)
					}
				}
			}
		}
	}

	result := 1
	myTicket := parseTicket(strings.Split(myTicketString, "\n")[1])
	for field, position := range fieldPositions {
		if strings.Contains(field, "departure") {
			result *= myTicket[position]
		}
	}
	return result
}

func parseInput(input string) (string, string, string) {
	split := strings.Split(input, "\n\n")
	return split[0], split[1], split[2]
}

func getRanges(rulesString string) map[string][][]int {
	ranges := make(map[string][][]int)
	for _, ruleString := range strings.Split(rulesString, "\n") {
		split := strings.Split(ruleString, ": ")
		field := split[0]
		rangesString := split[1]
		rangeStrings := strings.Split(rangesString, " or ")
		for _, rangeString := range rangeStrings {
			ranges[field] = append(ranges[field], MapToInts(strings.Split(rangeString, "-")))
		}
	}
	return ranges
}

func getNearbyTickets(nearbyTicketsString string) (nearbyTickets [][]int) {
	for i, nearbyTicketString := range strings.Split(nearbyTicketsString, "\n") {
		if i == 0 {
			continue
		}
		nearbyTickets = append(nearbyTickets, parseTicket(nearbyTicketString))
	}
	return
}

func parseTicket(ticket string) []int {
	return MapToInts(strings.Split(ticket, ","))
}

func getInvalidNearbyTickets(nearbyTickets [][]int, fields map[string][][]int) (invalid []int) {
	for _, ticket := range nearbyTickets {
		for _, value := range ticket {
			isInvalid := true
			for _, ranges := range fields {
				for _, r := range ranges {
					if isInRange(value, r) {
						isInvalid = false
						break
					}
				}
			}
			if isInvalid {
				invalid = append(invalid, value)
			}
		}
	}
	return
}

func removeTicketsContainingValues(tickets [][]int, values []int) (result [][]int) {
	for _, ticket := range tickets {
		include := true
		for _, y := range values {
			if contains(ticket, y) {
				include = false
			}
		}
		if include {
			result = append(result, ticket)
		}
	}
	return
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func isInRange(v int, r []int) bool {
	return v >= r[0] && v <= r[1]
}

func remove(slice []int, x int) (result []int) {
	for _, v := range slice {
		if v != x {
			result = append(result, v)
		}
	}
	return
}
