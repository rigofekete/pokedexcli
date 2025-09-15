package main

import (
		"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "  Hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "   hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input: " hello ",
			expected: []string{"hello"},
		},
	
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("lengths don't match: '%v' vs '%v'", actual, c.expected)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("actual word - %v -  does not match expected word - %v - ", word, expectedWord) 
			}
		}
	}
}


