package attempt1

import (
	"fmt"
	"io/ioutil"
	"testing"
	. "utils"
)

func TestPuzzle1(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			"../test-input-1.txt",
			2,
		},
		{
			"../test-input-2.txt",
			8,
		},
		{
			"../test-input-3.txt",
			2,
		},
	}

	for _, test := range tests {
		a := puzzle1(getInput(test.input))
		if a != test.expected {
			t.Errorf("expected: %d, actual: %d", test.expected, a)
		}
	}
}

func TestExpandBrackets(t *testing.T) {
	tests := []struct {
		brackets string
		expected string
	}{
		{
			"(a)((b)|(c))",
			"ab|ac",
		},
		{
			"(a)(bc)",
			"abc",
		},
		{
			"(a)(b|c)",
			"ab|ac",
		},
		{
			"(((a)))(b)",
			"ab",
		},
		{
			"(a|b)(a|b)",
			"aa|ab|ba|bb",
		},
		{
			"(a)(ab|ba)(b)",
			"aabb|abab",
		},
		{
			"(a|b|c)(a)",
			"aa|ba|ca",
		},
		{
			"(a|b|c)(a|b)",
			"aa|ab|ba|bb|ca|cb",
		},
		{
			"(a|b|c)(a|b|c)",
			"aa|ab|ac|ba|bb|bc|ca|cb|cc",
		},
		{
			"(a|b|c)(a|b|c)(a|b|c)",
			"aaa|aab|aac|aba|abb|abc|aca|acb|acc|baa|bab|bac|bba|bbb|bbc|bca|bcb|bcc|caa|cab|cac|cba|cbb|cbc|cca|ccb|ccc",
		},
		{
			"((a|b)|(a|b)|(a|b))(a|b)",
			"aa|ab|ba|bb|aa|ab|ba|bb|aa|ab|ba|bb",
		},
	}

	for _, test := range tests {
		a := expandBrackets(test.brackets)
		if a != test.expected {
			t.Errorf("expected: %s, actual: %s", test.expected, a)
		}
	}
}

func TestSolvePuzzle1(t *testing.T) {
	fmt.Println("Puzzle 1:", puzzle1(getInput("../input.txt")))
}

func getInput(path string) string {
	file, err := ioutil.ReadFile(path)
	Check(err)
	return string(file)
}
