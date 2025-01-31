package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " HellO World ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " HELLO ",
			expected: []string{"hello"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// check lenght of the actual slice
		if len(actual) != len(c.expected) {
			t.Errorf("For input '%v', lengths don't match: got '%v', want '%v'", c.input, actual, c.expected)
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("For input '%v', words don't match at index %d: got '%v', want '%v'", c.input, i, word, expectedWord)
			}
		}
	}
}
