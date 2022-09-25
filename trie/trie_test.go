package trie

import (
	"bufio"
	"math/rand"
	"os"
	"reflect"
	"sort"
	"testing"
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

func TestDeleteWord(t *testing.T) {
	for _, test := range deleteWordTestCases {
		testName := test.name
		testTrie := test.trie.deepCopy()
		testWord := test.input
		succeeds := test.succeeds
		expectedTrie := test.expected.deepCopy()

		t.Run(testName, func(t *testing.T) {
			ok := testTrie.DeleteWord(testWord)

			if ok != succeeds {
				if succeeds {
					t.Errorf("Expected DeleteWord(\"%s\") to succeed but it didn't find the word", testWord)
				} else {
					t.Errorf("Expected DeleteWord(\"%s\") to fail but it did find the word", testWord)
				}
			}

			if !reflect.DeepEqual(testTrie, expectedTrie) {
				t.Errorf("The trie after DeleteWord(\"%s\") doesn't match expected", testWord)
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
	for _, test := range getWordsFromPrefixTestCases {
		testName := test.name
		testTrie := test.trie
		testPrefix := test.prefix
		testExpected := test.expected

		sort.Strings(testExpected)

		t.Run(testName, func(t *testing.T) {
			results := testTrie.GetWordsFromPrefix(testPrefix)
			sort.Strings(results)

			if len(results) != len(testExpected) {
				t.Errorf("Expected GetWordsFromPrefix result to be %v but got %v", testExpected, results)
			}

			for i := range results {
				resultWord := results[i]
				expectedWord := testExpected[i]

				if resultWord != expectedWord {
					t.Errorf("Expected GetWordsFromPrefix result to be %v but got %v", testExpected, results)
					break
				}
			}
		})
	}

}

func BenchmarkGetWordsFromPrefix(b *testing.B) {
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

	trie := New()
	trie.AddWords(words)

	prefixes := make([]string, 0)
	for i := 0; i < b.N; i++ {
		randomWord := words[rand.Intn(len(words))]
		prefix := randomWord[:rand.Intn(len(randomWord))]
		prefixes = append(prefixes, prefix)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		trie.GetWordsFromPrefix(prefixes[i])
	}

	b.StopTimer()
}
