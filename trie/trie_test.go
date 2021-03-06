package trie

import (
	"testing"
	"reflect"
	"os"
	"bufio"
	"math/rand"
	"sort"
)

func TestAddWord(t *testing.T) {
	for _, test := range addWordTestCases {
		testName := test.name
		testWord := test.input
		expectedTrie := &test.expected

		t.Run(testName, func(t *testing.T) {
			trie := New()

			trie.AddWord(testWord)

			if !reflect.DeepEqual(trie, expectedTrie) {
				t.Errorf("The trie after adding word '%s' does not match the expected result", testWord)
			}
		})
	}
}

func BenchmarkAddWord(b *testing.B) {
	file, err := os.Open("../data/words.txt")
	if err != nil {
		b.Errorf("%s\nMake sure to unzip data/words.zip into data/words.txt before running benchmarks", err)
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
		testName := test.name
		testWords := test.input
		expectedTrie := &test.expected

		t.Run(testName, func(t *testing.T) {
			trie := New()

			trie.AddWords(testWords)

			if !reflect.DeepEqual(trie, expectedTrie) {
				t.Errorf("The trie after adding words %v does not match the expected result", testWords)
			}
		})
	}
}

func BenchmarkAddWords(b *testing.B) {
	file, err := os.Open("../data/words.txt")
	if err != nil {
		b.Errorf("%s\nMake sure to unzip data/words.zip into data/words.txt before running benchmarks", err)
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
		testName := test.name
		testTrie := test.trie
		testWord := test.input
		expectedEndNode := test.expectedEndNode
		expectedLeftover := test.expectedLeftover

		t.Run(testName, func(t *testing.T) {
			resultEndNode, resultLeftover := testTrie.traverseTrie(testWord)

			if resultEndNode != expectedEndNode {
				t.Errorf(
					"The returned node %v does not match the expected node %v",
					resultEndNode,
					expectedEndNode,
				)
				return
			}

			if resultLeftover != expectedLeftover {
				t.Errorf("The returned leftover string %s does not match the expected leftover string %s", resultLeftover, expectedLeftover)
			}
		})
	}
}

func TestHasWord(t *testing.T) {
	for _, test := range hasWordTestCases {
		testName := test.name
		testTrie := test.trie
		testWord := test.input
		testExpected := test.expected

		t.Run(testName, func(t *testing.T) {
			result := testTrie.HasWord(testWord)

			if result != testExpected {
				t.Errorf("Expected HasWord(\"%s\") to return %t but got %t", testWord, testExpected, result)
			}
		})
	}
}

func TestDFSWords(t *testing.T) {
	for _, test := range dfsWordsTestCases {
		testName := test.name
		testStartNode := test.startNode
		testPrefix := test.prefix
		testExpected := test.expected

		sort.Strings(testExpected)

		t.Run(testName, func(t *testing.T) {
			results := testStartNode.dfsWords(testPrefix)
			sort.Strings(results)

			if len(results) != len(testExpected) {
				t.Errorf("Expected dfsWords result to be %v but got %v", testExpected, results)
			}

			for i := range results {
				resultWord := results[i]
				expectedWord := testExpected[i]

				if resultWord != expectedWord {
					t.Errorf("Expected dfsWords result to be %v but got %v", testExpected, results)
					break
				}
			}
		})
	}
}

func TestGetWordsFromPrefix(t *testing.T) {
	t.Skip("not implemented")
}

func BenchmarkGetWordsFromPrefix(b *testing.B) {
	b.Skip("not implemented")
}
