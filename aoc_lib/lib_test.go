package aoc_lib

import "testing"

func TestInput(t *testing.T) {
	expected := "THIS\nis\nINPUT!"
	input := getInput("testing_resources/input")
	if string(input) != expected {
		t.Fatalf("Unmatched input\n%s\n%s", string(input), expected)
	}
}
