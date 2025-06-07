package repl

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
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

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("func cleanInput(); Length does not match. %v does not match %v!", len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("func cleanInput(); Word does not match %s does not match %s!", word, expectedWord)
			}
		}
	}
}
