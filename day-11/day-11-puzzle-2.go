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
	fmt.Println(occupiedSeats)
}

func applyRules(plan []string) []string {
	newPlan := []string{}
	for y, row := range plan {
		newRow := ""
		for x, l := range row {
			loc := string(l)

			if loc == "#" {
				if countVisibleOccupiedSeats(plan, x, y) >= 5 {
					newRow += "L"
					continue
				}
			}

			if loc == "L" {
				if countVisibleOccupiedSeats(plan, x, y) == 0 {
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

func countVisibleOccupiedSeats(plan []string, x int, y int) int {
	result := 0
	row := plan[y]
	for i := 1; x-i >= 0; i++ {
		s := string(row[x-i])
		if s == "L" {
			break
		}

		if s == "#" {
			result++
			break
		}
	}

	for i := 1; x+i <= len(row)-1; i++ {
		s := string(row[x+i])
		if s == "L" {
			break
		}

		if s == "#" {
			result++
			break
		}
	}

	for j := 1; y-j >= 0; j++ {
		prevRow := plan[y-j]
		s := string(prevRow[x])
		if s == "L" {
			break
		}

		if s == "#" {
			result++
			break
		}
	}

	for j := 1; y-j >= 0; j++ {
		prevRow := plan[y-j]
		if x-j < 0 {
			break
		}
		s := string(prevRow[x-j])

		if s == "L" {
			break
		}

		if s == "#" {
			result++
			break
		}
	}

	for j := 1; y-j >= 0; j++ {
		prevRow := plan[y-j]
		if x+j > len(row)-1 {
			break
		}
		s := string(prevRow[x+j])

		if s == "L" {
			break
		}

		if s == "#" {
			result++
			break
		}
	}

	for j := 1; y+j <= len(plan)-1; j++ {
		nextRow := plan[y+j]
		s := string(nextRow[x])
		if s == "L" {
			break
		}

		if s == "#" {
			result++
			break
		}
	}

	for j := 1; y+j <= len(plan)-1; j++ {
		nextRow := plan[y+j]
		if x-j < 0 {
			break
		}
		s := string(nextRow[x-j])
		if s == "L" {
			break
		}

		if s == "#" {
			result++
			break
		}
	}

	for j := 1; y+j <= len(plan)-1; j++ {
		nextRow := plan[y+j]
		if x+j > len(row)-1 {
			break
		}
		s := string(nextRow[x+j])
		if s == "L" {
			break
		}

		if s == "#" {
			result++
			break
		}
	}

	return result
}

func countOccupiedSeats(seats string) int {
	r := regexp.MustCompile("#")
	return len(r.FindAllString(seats, -1))
}
