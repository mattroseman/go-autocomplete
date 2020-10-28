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

		trie := NewTrie()

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

	trie := NewTrie()
	for i := 0; i < b.N; i++ {
		trie.AddWord(words[rand.Intn(len(words))])
	}

	b.StopTimer()
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
		trie := NewTrie()
		trie.AddWords(words)
	}

	b.StopTimer()
}
