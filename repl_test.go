package main

import (
	"testing"
	"fmt"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello world   ",
			expected: []string{"hello", "world"},
		},
		{
			input: "HELLo   BelLa   ",
			expected: []string{"hello", "bella"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("len actual does not match len expected")
			fmt.Printf("actual: %s -- expected: %s", actual, c.expected)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("expected %q but got %q", expectedWord, word)
			}
		}
	}
}