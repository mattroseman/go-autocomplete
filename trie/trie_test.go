package trie

import (
	"testing"
	"reflect"
	"os"
	"bufio"
	"math/rand"
)

func TestAddWord(t *testing.T) {
	for _, test := range addWordTestCases {
		testWord := test.input
		expectedTrie := &test.expected

		trie := New()

		trie.AddWord(testWord)

		if !reflect.DeepEqual(trie, expectedTrie) {
			t.Errorf("The trie after adding word '%s' does not match the expected result", testWord)
		}
	}
}

func BenchmarkAddWord(b *testing.B) {
	file, err := os.Open("../data/words.txt")
	if err != nil {
		b.Error(err)
		return
	}
	defer file.Close()

	var words []string

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	b.ResetTimer()

	trie := New()
	for i := 0; i < b.N; i++ {
		trie.AddWord(words[rand.Intn(len(words))])
	}

	b.StopTimer()
}

func TestAddWords(t *testing.T) {
	for _, test := range addWordsTestCases {
		testWords := test.input
		expectedTrie := &test.expected

		trie := New()

		trie.AddWords(testWords)

		if !reflect.DeepEqual(trie, expectedTrie) {
			t.Errorf("The trie after adding words %v does not match the expected result", testWords)
		}
	}
}

func BenchmarkAddWords(b *testing.B) {
	file, err := os.Open("../data/words.txt")
	if err != nil {
		b.Error(err)
		return
	}
	defer file.Close()

	var words []string

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		trie := New()
		trie.AddWords(words)
	}

	b.StopTimer()
}

func TestTraverseTrie(t *testing.T) {
	for _, test := range traverseTrieTestCases {
		testTrie := test.trie
		testWord := test.input
		testSucceeds := test.succeeds
		testExpected := test.expected

		result, ok := testTrie.traverseTrie(testWord)

		if ok != testSucceeds {
			t.Errorf("The word %s should have been found", testWord)
			return
		}

		if result != testExpected {
			t.Errorf("the returned node %v does not match the expected node %v", *result, *testExpected)
		}
	}
}

func TestHasWord(t *testing.T) {
	for _, test := range hasWordTestCases {
		testTrie := test.trie
		testWord := test.input
		testExpected := test.expected

		result := testTrie.HasWord(testWord)

		if result != testExpected {
			t.Errorf("Expected HasWord(\"%s\") to return %t but got %t", testWord, testExpected, result)
		}
	}
}

func TestDFSWords(t *testing.T) {
	t.Skip("not implemented")
}

func TestGetWordsFromPrefix(t *testing.T) {
	t.Skip("not implemented")
}

func BenchmarkGetWordsFromPrefix(b *testing.B) {
	b.Skip("not implemented")
}
