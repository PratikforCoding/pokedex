package main

import (
	"testing"
	"fmt"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "  ",
			expected: []string{},
		},
		{
			input: "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input: "Hello World",
			expected: []string{"hello", "world"},
		},
	
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			fmt.Printf("Legth does not match: %s vs %s", actual, c.expected)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("cleanInput(%v) == %v, expected %v", c.input, actual, c.expected)
			}
		}
	}
}