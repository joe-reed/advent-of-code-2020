package main

import (
	"sort"
	"strings"
)

func puzzle1(input []string) (result int) {
	tiles := parseInput(input)
	tiles = simplifyTiles(tiles)
	tileStrings := []string{}
	for _, tile := range tiles {
		sort.Strings(tile)
		tileStrings = append(tileStrings, strings.Join(tile, ""))
	}
	counts := make(map[string]int)
	for _, tileString := range tileStrings {
		_, ok := counts[tileString]
		if !ok {
			counts[tileString] = 1
			continue
		}
		counts[tileString]++
	}
	for _, count := range counts {
		if count%2 == 1 {
			result++
		}
	}
	return
}

func puzzle2(input []string, rounds int) (result int) {
	tiles := parseInput(input)
	tiles = simplifyTiles(tiles)
	tileStrings := []string{}
	for _, tile := range tiles {
		tileStrings = append(tileStrings, getTileString(tile))
	}

	colours := make(map[string]string)
	for _, tileString := range tileStrings {
		colours[tileString] = "w"
	}

	colours = flipTiles(tileStrings, colours)

	tiles = unique(tiles)

	seen := map[string]bool{}
	for _, tile := range tiles {
		seen[getTileString(tile)] = true
	}

	for i := 0; i < rounds; i++ {
		for _, tile := range tiles {
			for _, adjacent := range getAdjacents(tile) {
				adjacentString := getTileString(adjacent)
				if _, a := seen[adjacentString]; !a {
					tiles = append(tiles, adjacent)
					seen[adjacentString] = true
					colours[adjacentString] = "w"
				}
			}
		}

		shouldFlip := []string{}
		for _, tile := range tiles {
			tileString := getTileString(tile)
			blackCount := 0
			for _, adjacent := range getAdjacents(tile) {
				if colours[getTileString(adjacent)] == "b" {
					blackCount++
				}
			}
			if colours[tileString] == "w" && blackCount == 2 {
				shouldFlip = append(shouldFlip, tileString)
				continue
			}
			if colours[tileString] == "b" && (blackCount == 0 || blackCount > 2) {
				shouldFlip = append(shouldFlip, tileString)
				continue
			}
		}

		colours = flipTiles(shouldFlip, colours)
	}

	for _, colour := range colours {
		if colour == "b" {
			result++
		}
	}
	return
}

func parseInput(input []string) (result [][]string) {
	for _, line := range input {
		result = append(result, parseLine(line))
	}
	return
}

func parseLine(line string) (result []string) {
	for len(line) > 0 {
		switch string(line[0]) {
		case "e", "w":
			result = append(result, string(line[0]))
			line = line[1:]
		case "n", "s":
			result = append(result, string(line[0:2]))
			line = line[2:]
		}
	}
	return
}

func simplifyTiles(tiles [][]string) (result [][]string) {
	for _, tile := range tiles {
		result = append(result, simplify(tile))
	}
	return
}

func simplify(tile []string) []string {
	result := make([]string, len(tile))
	copy(result, tile)
	for {
		if contains(result, "e") && contains(result, "w") {
			result = removeOne(result, "e")
			result = removeOne(result, "w")
			continue
		}
		if contains(result, "nw") && contains(result, "se") {
			result = removeOne(result, "nw")
			result = removeOne(result, "se")
			continue
		}
		if contains(result, "ne") && contains(result, "sw") {
			result = removeOne(result, "ne")
			result = removeOne(result, "sw")
			continue
		}
		if contains(result, "nw") && contains(result, "sw") {
			result = removeOne(result, "nw")
			result = removeOne(result, "sw")
			result = append(result, "w")
			continue
		}
		if contains(result, "ne") && contains(result, "se") {
			result = removeOne(result, "ne")
			result = removeOne(result, "se")
			result = append(result, "e")
			continue
		}
		if contains(result, "e") && contains(result, "nw") {
			result = removeOne(result, "e")
			result = removeOne(result, "nw")
			result = append(result, "ne")
			continue
		}
		if contains(result, "w") && contains(result, "ne") {
			result = removeOne(result, "w")
			result = removeOne(result, "ne")
			result = append(result, "nw")
			continue
		}
		if contains(result, "e") && contains(result, "sw") {
			result = removeOne(result, "e")
			result = removeOne(result, "sw")
			result = append(result, "se")
			continue
		}
		if contains(result, "w") && contains(result, "se") {
			result = removeOne(result, "w")
			result = removeOne(result, "se")
			result = append(result, "sw")
			continue
		}
		break
	}
	return result
}

func contains(s []string, e string) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func removeOne(s []string, e string) []string {
	result := make([]string, len(s)-1)
	for i, v := range s {
		if v != e {
			continue
		}
		copy(result[:i], s[:i])
		if i < len(s)-1 {
			copy(result[i:], s[i+1:])
		}
		break
	}
	return result
}

func flipTiles(tilesToFlip []string, colours map[string]string) map[string]string {
	result := make(map[string]string)
	for k, v := range colours {
		result[k] = v
	}
	for _, tileString := range tilesToFlip {
		if result[tileString] == "w" {
			result[tileString] = "b"
			continue
		}
		result[tileString] = "w"
	}
	return result
}

func getTileString(tile []string) string {
	c := make([]string, len(tile))
	copy(c, tile)
	sort.Strings(c)
	return strings.Join(c, "")
}

func unique(slice [][]string) (result [][]string) {
	keys := make(map[string]bool)
	for _, tile := range slice {
		tileString := getTileString(tile)
		if _, value := keys[tileString]; !value {
			keys[tileString] = true
			result = append(result, tile)
		}
	}
	return
}

var adjacentCache = map[string][][]string{}

func getAdjacents(tile []string) (result [][]string) {
	tileString := getTileString(tile)
	cachedResult, ok := adjacentCache[tileString]

	if ok {
		return cachedResult
	}

	for _, direction := range []string{"e", "se", "sw", "w", "nw", "ne"} {
		adjacent := simplify(append(tile, direction))
		result = append(result, adjacent)
	}
	adjacentCache[tileString] = result
	return
}
