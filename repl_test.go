package main

import (
"testing"
)

func TestCleanInput(t *testing.T) {
    cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "My favorite game is Bioshock !",
			expected: []string{"my", "favorite", "game", "is", "bioshock", "!"},
		},
		{
			input: " fourty       three ",
			expected: []string{"fourty", "three"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Length of strings do not match")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("For input '%s': at index %d, expected '%s' but got '%s'", c.input, i, expectedWord, word)
			}
		}
	}
}