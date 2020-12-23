package main

import "container/ring"

func puzzle1(input []int, moves int) int {
	cups := ring.New(len(input))
	for _, cup := range input {
		cups.Value = cup
		cups = cups.Next()
	}
	for i := 0; i < moves; i++ {
		cups = performMove(cups)
	}
	return getAnswer(cups)
}

func performMove(cups *ring.Ring) *ring.Ring {
	current := cups.Value.(int)
	pickedUp := cups.Unlink(3)
	destination := current - 1
	for cups.Value != destination {
		cups = cups.Next()
		if cups.Value == current {
			destination--
		}
		if destination < min(cups) {
			destination = max(cups)
		}
	}
	cups.Link(pickedUp)
	for cups.Value != current {
		cups = cups.Next()
	}

	return cups.Next()
}

func getAnswer(cups *ring.Ring) (result int) {
	for cups.Value != 1 {
		cups = cups.Next()
	}
	cups.Next()
	cups.Do(func(cup interface{}) {
		if cup.(int) != 1 {
			result = result*10 + cup.(int)
		}
	})
	return
}

func min(cups *ring.Ring) int {
	var min = cups.Value.(int)
	cups.Do(func(cup interface{}) {
		if cup.(int) < min {
			min = cup.(int)
		}
	})
	return min
}

func max(cups *ring.Ring) (max int) {
	cups.Do(func(cup interface{}) {
		if cup.(int) > max {
			max = cup.(int)
		}
	})
	return
}
