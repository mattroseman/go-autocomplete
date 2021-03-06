// Package trie implements a minified Deterministic Acyclic Finite State Automata (DAFSA).
// https://en.wikipedia.org/wiki/Deterministic_acyclic_finite_state_automaton
// A DAFSA data structure is useful for getting possible given a certain prefix which in turn can be used for autocomplete.
package trie

import "strings"

type node struct {
	edgeLabel string
	endOfWord bool
	children map[byte]*node // maps a single character to a child node that will contain the rest of the edgeLabel
}

type Trie struct {
	root *node
}

// New constructs a new trie instance with a root node with no children.
func New() *Trie {
	return &Trie{
		root: &node{
			edgeLabel: "",
			endOfWord: false,
			children: make(map[byte]*node),
		},
	}
}

// AddWord adds a given word to the trie t.
// It returns true if a new word was added or false if the word already existed.
func (t *Trie) AddWord(word string) bool {
	if len(word) == 0 {
		return false
	}

	word = strings.ToLower(word)

	currentNode := t.root
	newWordAdded := false

	for i := 0; i < len(word); {
		char := word[i]

		// if a child node exists for char
		if child, ok := currentNode.children[char]; ok {
			// if the edge label for the child node matches what's left of the word
			// edge label: 'abc', what's left of the word: 'abc'
			if child.edgeLabel == word[i:] {
				newWordAdded = !child.endOfWord

				// mark the child node as an end of a word
				child.endOfWord = true

				break
			}

			commonPrefix := getCommonPrefix(child.edgeLabel, word[i:])

			// if the edge label for the child contains the entirety of what's left of the word plus some extra
			// edge label: 'abcdef', what's left of the word: 'abc'
			if len(commonPrefix) < len(child.edgeLabel) && len(commonPrefix) == len(word[i:]) {
				// insert a new node between current node and child

				// make the new node with an edge label of what's left of the word and add child to children
				newNode := &node{
					edgeLabel: word[i:],
					endOfWord: true,
					children: map[byte]*node{child.edgeLabel[len(word[i:])]: child},
				}

				// change the child's edge label to be the leftover of the previous edge label
				child.edgeLabel = child.edgeLabel[len(word[i:]):]

				// make the new node a child of the current node
				currentNode.children[char] = newNode

				newWordAdded = true

				break
			}

			// if the edge label and what's left of the word share a common prefix, but differ at some point
			// edge label: 'abcdef', what's left of the word: 'abcxyz'
			if len(commonPrefix) < len(child.edgeLabel) && len(commonPrefix) < len(word[i:]) {
				// insert a branching node between current node and child and give it a new node child

				// create a new node that will be a child of currentNode and have two edges one going to child and one to another new node
				newNode := &node{
					edgeLabel: commonPrefix,
					endOfWord: false,
					children: map[byte]*node{child.edgeLabel[len(commonPrefix)]: child},
				}

				// change the child's edge label to be the leftover of the previous edge label
				child.edgeLabel = child.edgeLabel[len(commonPrefix):]

				// make the new node a child of the current node
				currentNode.children[char] = newNode

				// add another new node as a child of the previous new node with a label of what's leftover of the word
				newNode.children[word[len(commonPrefix)]] = &node{
					edgeLabel: word[len(commonPrefix):],
					endOfWord: true,
					children: make(map[byte]*node),
				}

				newWordAdded = true

				break
			}

			// the last option is that what's left of the word contains the entirety of the edge label plus some extra
			// edge label: 'abc', what's left of the word: 'abcdef'
			// follow the edge to the next node and increment i to account for the edge's label
			i += len(child.edgeLabel)
			currentNode = child
		} else {
			currentNode.children[char] = &node{
				edgeLabel: word[i:],
				endOfWord: true,
				children: make(map[byte]*node),
			}

			newWordAdded = true

			break
		}
	}

	return newWordAdded
}

// AddWords adds the given array of words to the trie t.
// It returns a uint32 indicating the number of new words that where added, ignoring duplicates.
func (t *Trie) AddWords(words []string) uint32 {
	var wordsAdded uint32 = 0

	for i := 0; i < len(words); i++ {
		if t.AddWord(words[i]) {
			wordsAdded++
		}
	}

	return wordsAdded
}

// DeleteWord removes the given word from the trie t.
// It returns a boolean of true if the given word was found and removed, false if the given word couldn't be found.
func (t *Trie) DeleteWord(word string) bool {
	// TODO implement
	return false
}

// traverseTrie traverses down the trie with the given word
// It returns the final node found or nil if nothing is found
func (t *Trie) traverseTrie(word string) (*node, bool) {
	word = strings.ToLower(word)
	currentNode := t.root

	for i := 0; i < len(word); {
		char := word[i]

		if child, ok := currentNode.children[char]; ok {
			commonPrefix := getCommonPrefix(child.edgeLabel, word[i:])

			// if the child's edge label differs from the given word, the given word doesn't exist in the trie
			// this can be tested if the common prefix is different than the current edge label
			if (commonPrefix != child.edgeLabel) {
				break
			}

			// if the common prefix matches what's left of the word, than the child is the final node
			if (commonPrefix == word[i:] && child.endOfWord) {
				return child, true
			}

			i += len(child.edgeLabel)
			currentNode = child
		} else {
			break
		}
	}

	return nil, false
}

// HasWord checks to see if the given word exists in the trie t.
// It returns a boolean of true if the given word exists in the trie, false otherwise.
func (t Trie) HasWord(word string) bool {
	_, ok := t.traverseTrie(word)

	return ok
}

// dfsTrie depth first searches the trie starting at the given node.
// It returns an array of words that are found from endOfWord node children.
func (start *node) dfsWords(prefix string) []string {
	words := make([]string, 0)

	if start.endOfWord {
		words = append(words, prefix)
	}

	for _, child := range start.children {
		words = append(words, child.dfsWords(prefix + child.edgeLabel)...)
	}

	return words
}

// GetWordsFromPrefix returns all the words in trie t that have the given prefix.
func (t Trie) GetWordsFromPrefix(prefix string) []string {
	words := make([]string, 0)

	prefixEndNode, ok := t.traverseTrie(prefix)
	// if the prefix isn't even found in the trie, return the empty array
	if !ok {
		return words
	}

	return prefixEndNode.dfsWords(prefix)
}
