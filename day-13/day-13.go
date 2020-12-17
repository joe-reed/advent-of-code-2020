package main

import (
	"math/big"
	"strconv"
	"strings"
)

func puzzle1(startTime int, schedule string) int {
	i := startTime
	for true {
		for _, busId := range getBusIds(schedule) {
			if i%busId == 0 {
				return busId * (i - startTime)
			}
		}
		i++
	}
	return 0
}

func getBusIds(schedule string) []int {
	ss := strings.Split(schedule, ",")
	var ids []int
	for _, s := range ss {
		id, err := strconv.Atoi(s)
		if err == nil {
			ids = append(ids, id)
		}
	}
	return ids
}

func puzzle2(schedule string) int {
	idRemainderMap := getIdRemainderMap(schedule)
	prod := 1
	for id, _ := range idRemainderMap {
		prod *= id
	}

	r := 0
	for id, rem := range idRemainderMap {
		r += rem * prod / id * modInverse(prod/id, id)
	}
	return r % prod
}

func getIdRemainderMap(schedule string) map[int]int {
	ss := strings.Split(schedule, ",")
	r := make(map[int]int)
	for i, s := range ss {
		id, err := strconv.Atoi(s)
		if err == nil {
			var rem int
			rem = (id - i) % id
			for rem < 0 {
				rem += id
			}
			r[id] = rem
		}
	}
	return r
}

func modInverse(a, b int) int {
	x := big.NewInt(int64(a))
	y := big.NewInt(int64(b))
	r := new(big.Int)
	r.ModInverse(x, y)
	return int(r.Int64())
}
