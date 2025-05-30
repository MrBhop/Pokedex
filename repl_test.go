package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	fmt.Println("Starting testing")
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "hello WORLD",
			expected: []string{"hello", "world"},
		},
		{
			input: "",
			expected: []string{},
		},
		{
			input: "   ",
			expected: []string{},
		},
		{
			input: "  wOrd 	 ",
			expected: []string{"word"},
		},
		{
			input: "i am a STRING of  text   ",
			expected: []string{"i", "am", "a", "string", "of", "text"},
		},
	}

	fmt.Printf("Number of test cases: %v\n", len(cases))
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Length does not match. %v does not match %v!", len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				fmt.Println("Failed Test")
				t.Errorf("Word does not match %s does not match %s!", word, expectedWord)
			}
		}
	}
}
