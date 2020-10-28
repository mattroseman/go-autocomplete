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

		err := trie.AddWord(testWord)
		if err != nil {
			t.Errorf("Something went wrong adding word '%s' to the trie\n%s\n", testWord, err)
		}

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

		err := trie.AddWords(testWords)
		if err != nil {
			t.Errorf("Something went wrong adding words %v to the trie\n%s\n", testWords, err)
		}

		if !reflect.DeepEqual(trie, expectedTrie) {
			t.Errorf("The trie after adding words %v does not match the expected result", testWords)
		}
	}
}
