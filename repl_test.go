package main

import "testing"

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
			input:    "  helloworld  ",
			expected: []string{"helloworld"},
		},
		{
			input:    "  hello  WORLD  ",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("Incorrect number of words. Input string was not split correctly.\n Actual: %v\nExpected: %v", actual, c.expected)
		} else {
			for i := range actual {
				word := actual[i]
				expectedWord := c.expected[i]

				if word != expectedWord {
					t.Errorf("Incorrect word. Expecting %v but found %v", expectedWord, word)
				}
			}
		}
	}
}
