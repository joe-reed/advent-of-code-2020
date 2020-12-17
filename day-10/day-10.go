package main

import (
	"sort"
)

func puzzle1(jolts []int) int {
	sort.Ints(jolts)

	jolts = append(jolts, jolts[len(jolts)-1]+3)
	differences := make(map[int]int)
	for i, jolt := range jolts {
		var diff int
		if i == 0 {
			diff = jolt
		} else {
			diff = jolt - jolts[i-1]
		}

		if _, ok := differences[diff]; ok {
			differences[diff]++
		} else {
			differences[diff] = 1
		}
	}

	return differences[1] * differences[3]
}

func puzzle2(jolts []int) int {
	sort.Ints(jolts)

	deviceJolts := jolts[len(jolts)-1] + 3
	jolts = append(jolts, deviceJolts)
	jolts = append([]int{0}, jolts...)

	branches := map[int]int{deviceJolts: 1}

	return countBranches(jolts, &branches)
}

func countBranches(tree []int, branches *map[int]int) int {
	count := 0
	i := 1
	option := tree[1]
	for option <= tree[0]+3 {
		if _, ok := (*branches)[option]; !ok {
			(*branches)[option] = countBranches(tree[i:], branches)
		}

		count += (*branches)[option]

		i++
		if i == len(tree) {
			break
		}
		option = tree[i]
	}
	return count
}
