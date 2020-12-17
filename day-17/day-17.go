package main

import (
	"strings"
	"sync"
)

func puzzle1(input string) int {
	activeCoords := getActiveCoords(input, 3)

	for i := 0; i < 6; i++ {
		prevActiveCoords := activeCoords
		activeCoords = [][]int{}
		start, end := getGridToConsider(prevActiveCoords, 3)
		for x := start[0]; x <= end[0]; x++ {
			for y := start[1]; y <= end[1]; y++ {
				for z := start[2]; z <= end[2]; z++ {
					coord := []int{x, y, z}
					if shouldBeActive(coord, prevActiveCoords, 3) {
						activeCoords = append(activeCoords, coord)
					}
				}
			}
		}
	}
	return len(activeCoords)
}

func puzzle2(input string) int {
	activeCoords := getActiveCoords(input, 4)

	for i := 0; i < 6; i++ {
		prevActiveCoords := activeCoords
		activeCoords = [][]int{}
		start, end := getGridToConsider(prevActiveCoords, 4)
		for x := start[0]; x <= end[0]; x++ {
			for y := start[1]; y <= end[1]; y++ {
				for z := start[2]; z <= end[2]; z++ {
					for w := start[3]; w <= end[3]; w++ {
						coord := []int{x, y, z, w}
						if shouldBeActive(coord, prevActiveCoords, 4) {
							activeCoords = append(activeCoords, coord)
						}
					}
				}
			}
		}
	}
	return len(activeCoords)
}

func puzzle2Concurrent(input string) int {
	activeCoords := getActiveCoords(input, 4)

	for i := 0; i < 6; i++ {
		prevActiveCoords := activeCoords
		start, end := getGridToConsider(prevActiveCoords, 4)
		var wg sync.WaitGroup
		c := make(chan []int)

		for x := start[0]; x <= end[0]; x++ {
			wg.Add(1)
			go func(x int) {
				defer wg.Done()
				for y := start[1]; y <= end[1]; y++ {
					for z := start[2]; z <= end[2]; z++ {
						for w := start[3]; w <= end[3]; w++ {
							coord := []int{x, y, z, w}
							if shouldBeActive(coord, prevActiveCoords, 4) {
								c <- coord
							}
						}
					}
				}
			}(x)
		}
		go func() {
			wg.Wait()
			close(c)
		}()

		activeCoords = [][]int{}
		for activeCoord := range c {
			activeCoords = append(activeCoords, activeCoord)
		}
	}
	return len(activeCoords)
}

func getActiveCoords(input string, dimensions int) (coords [][]int) {
	lines := strings.Split(input, "\n")
	for x := 0; x < len(lines[0]); x++ {
		for y := 0; y < len(lines); y++ {
			if string(lines[x][y]) == "#" {
				coord := []int{x, y}
				for len(coord) < dimensions {
					coord = append(coord, 0)
				}
				coords = append(coords, coord)
			}
		}
	}
	return
}

func getGridToConsider(coords [][]int, dimensions int) (start, end []int) {
	for i := 0; i < dimensions; i++ {
		min, max := 0, 0
		for _, coord := range coords {
			if coord[i] < min {
				min = coord[i]
			}
			if coord[i] > max {
				max = coord[i]
			}
		}
		start, end = append(start, min-1), append(end, max+1)
	}
	return
}

func shouldBeActive(coord []int, prevActiveCoords [][]int, dimensions int) bool {
	count := countActiveNeighbours(prevActiveCoords, coord, dimensions)
	if contains(prevActiveCoords, coord, dimensions) {
		return count == 2 || count == 3
	}
	return count == 3
}

func countActiveNeighbours(activeCoords [][]int, coord []int, dimensions int) (count int) {
	for _, activeCoord := range activeCoords {
		if equals(activeCoord, coord, dimensions) {
			continue
		}
		if isNeighbour(coord, activeCoord, dimensions) {
			count++
		}
	}
	return
}

func isNeighbour(a []int, b []int, dimensions int) bool {
	for i := 0; i < dimensions; i++ {
		d := a[i] - b[i]
		if d < 0 {
			d *= -1
		}
		if d > 1 {
			return false
		}
	}
	return true
}

func contains(coords [][]int, coord []int, dimensions int) bool {
	for _, c := range coords {
		if equals(c, coord, dimensions) {
			return true
		}
	}
	return false
}

func equals(a []int, b []int, dimensions int) bool {
	for i := 0; i < dimensions; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
