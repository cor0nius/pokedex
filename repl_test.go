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
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "HeLlO\nWoRlD",
			expected: []string{"hello", "world"},
		},
		{
			input:    "This String Is A Bit Longer",
			expected: []string{"this", "string", "is", "a", "bit", "longer"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("FAIL - different slice length")
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("FAIL - words at index %v do not match", i)
				break
			}
		}
	}
}
