package trie

import (
	"testing"
	"reflect"
)

func TestAddWord(t *testing.T) {
	for _, test := range addWordTestCases {
		testWord := test.input
		expectedTrie := &test.expected

		trie := NewTrie()

		trie.AddWord(testWord)

		if !reflect.DeepEqual(trie, expectedTrie) {
			t.Errorf("The trie after adding word '%s' does not match the expected result", testWord)
		}
	}
}

func TestAddWords(t *testing.T) {
	for _, test := range addWordsTestCases {
		testWords := test.input
		expectedTrie := &test.expected

		trie := NewTrie()

		trie.AddWords(testWords)

		if !reflect.DeepEqual(trie, expectedTrie) {
			t.Errorf("The trie after adding words %v does not match the expected result", testWords)
		}
	}
}
