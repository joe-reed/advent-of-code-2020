package day11

import (
	"fmt"
	"reflect"
	"regexp"
	. "utils"
)

func main() {
	plan := LoadFile("./input.txt")
	newPlan := []string{}
	for !reflect.DeepEqual(plan, newPlan) {
		newPlan = plan
		plan = applyRules(newPlan)
	}

	occupiedSeats := 0
	for _, row := range plan {
		occupiedSeats += countOccupiedSeats(row)
	}
	fmt.Println(plan)
	fmt.Println(occupiedSeats)
}

func applyRules(plan []string) []string {
	newPlan := []string{}
	for y, row := range plan {
		newRow := ""
		for x, l := range row {
			loc := string(l)

			if loc == "#" {
				if countOccupiedSeats(getAdjacentSeats(plan, x, y)) >= 4 {
					newRow += "L"
					continue
				}
			}

			if loc == "L" {
				if countOccupiedSeats(getAdjacentSeats(plan, x, y)) == 0 {
					newRow += "#"
					continue
				}
			}

			newRow += loc
		}
		newPlan = append(newPlan, newRow)
	}
	return newPlan
}

func getAdjacentSeats(plan []string, x int, y int) string {
	result := ""
	row := plan[y]
	if x != 0 {
		result += string(row[x-1])
	}
	if x != len(row)-1 {
		result += string(row[x+1])
	}

	if y != 0 {
		prevRow := plan[y-1]
		result += string(prevRow[x])
		if x != 0 {
			result += string(prevRow[x-1])
		}
		if x != len(row)-1 {
			result += string(prevRow[x+1])
		}
	}

	if y != len(plan)-1 {
		nextRow := plan[y+1]
		result += string(nextRow[x])
		if x != 0 {
			result += string(nextRow[x-1])
		}
		if x != len(row)-1 {
			result += string(nextRow[x+1])
		}
	}

	return result
}

func countOccupiedSeats(seats string) int {
	r := regexp.MustCompile("#")
	return len(r.FindAllString(seats, -1))
}
