package main

import (
	"fmt"
	"regexp"
	"strings"
	. "utils"
)

func main() {
	input := LoadFile("./input.txt")

	var instructions []Instruction
	for _, line := range input {
		instructions = append(instructions, getInstruction(line))
	}

	acc, loc := 0, 0
	visited := []int{}

	for !IntInArray(loc, visited) {
		visited = append(visited, loc)
		acc, loc = instructions[loc].Execute(acc, loc)
	}
	fmt.Println(acc)
}

func getInstruction(line string) Instruction {
	re := regexp.MustCompile(` (.{1})(\d+)`)
	matches := re.FindStringSubmatch(line)

	if strings.Contains(line, "acc") {
		return Acc{direction: matches[1], amount: ConvertToInt(matches[2])}
	}

	if strings.Contains(line, "jmp") {
		return Jmp{direction: matches[1], amount: ConvertToInt(matches[2])}
	}

	return Nop{}
}

type Instruction interface {
	Execute(acc int, loc int) (int, int)
}

type Acc struct {
	direction string
	amount    int
}

func (a Acc) Execute(acc int, loc int) (int, int) {
	if a.direction == "+" {
		acc += a.amount
	} else {
		acc -= a.amount
	}
	loc += 1
	return acc, loc
}

type Jmp struct {
	direction string
	amount    int
}

func (l Jmp) Execute(acc int, loc int) (int, int) {
	if l.direction == "+" {
		loc += l.amount
	} else {
		loc -= l.amount
	}
	return acc, loc
}

type Nop struct{}

func (n Nop) Execute(acc int, loc int) (int, int) {
	loc += 1
	return acc, loc
}
