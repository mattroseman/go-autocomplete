package trie

import (
	"testing"
)

func TestGetCommonPrefix(t *testing.T) {
	tests := []struct{
		inputA string
		inputB string
		expected string
	}{
		{"abcdefg", "abcxyz", "abc"},
		{"banana", "cucumber", ""},
		{"bread", "bread", "bread"},
		{"a", "a", "a"},
		{"", "", ""},
		{"optics", "optimize", "opti"},
	}

	for _, test := range tests {
		prefix := getCommonPrefix(test.inputA, test.inputB)

		if prefix != test.expected {
			t.Errorf("Returned prefix '%s' doesn't match expected '%s' for strings '%s' and '%s'\n", prefix, test.expected, test.inputA, test.inputB)
		}
	}
}
