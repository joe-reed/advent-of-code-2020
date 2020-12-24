package main

import "container/ring"

func puzzle1(input []int, moves int) int {
	cups, cupsMap := parseInput(input)
	len := len(input)
	for i := 0; i < moves; i++ {
		pickedUp := cups.Unlink(3)

		destination := cups.Value.(int) - 1
		if destination == 0 {
			destination = len
		}
		for contains(pickedUp, destination) {
			destination--
			if destination == 0 {
				destination = len
			}
		}

		cupsMap[destination].Link(pickedUp)

		cups = cups.Next()
	}

	cups = cupsMap[1].Next()
	result := 0
	cups.Do(func(cup interface{}) {
		if cup.(int) != 1 {
			result = result*10 + cup.(int)
		}
	})
	return result
}

func puzzle2(input []int) int {
	cups, cupsMap := parseInput(input)
	remainingCups, cupsMap := getRemainingCups(input, cupsMap)
	cups.Prev().Link(remainingCups)

	len := cups.Len()
	for i := 0; i < 10000000; i++ {
		pickedUp := cups.Unlink(3)

		destination := cups.Value.(int) - 1
		if destination == 0 {
			destination = len
		}
		for contains(pickedUp, destination) {
			destination--
			if destination == 0 {
				destination = len
			}
		}

		cupsMap[destination].Link(pickedUp)

		cups = cups.Next()
	}

	cup1 := cupsMap[1]
	return cup1.Move(1).Value.(int) * cup1.Move(2).Value.(int)
}

func parseInput(input []int) (cups *ring.Ring, cupsMap map[int]*ring.Ring) {
	cups = ring.New(len(input))
	cupsMap = make(map[int]*ring.Ring)
	for _, cup := range input {
		cups.Value = cup
		cupsMap[cup] = cups
		cups = cups.Next()
	}
	return
}

func getRemainingCups(input []int, cupsMap map[int]*ring.Ring) (*ring.Ring, map[int]*ring.Ring) {
	size := 1000000
	cups := ring.New(size - len(input))
	for i := len(input) + 1; i <= size; i++ {
		cups.Value = i
		cupsMap[i] = cups
		cups = cups.Next()
	}
	return cups, cupsMap
}

func contains(cups *ring.Ring, c int) bool {
	result := false
	cups.Do(func(cup interface{}) {
		if cup.(int) == c {
			result = true
			return
		}
	})
	return result
}
