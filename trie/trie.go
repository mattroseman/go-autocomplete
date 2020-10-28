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

type trie struct {
	root *node
}

func NewTrie() *trie {
	return &trie{
		root: &node{
			edgeLabel: "",
			endOfWord: false,
			children: make(map[byte]*node),
		},
	}
}

// AddWord adds a given word to the trie t
func (t *trie) AddWord(word string) {
	if len(word) == 0 {
		return
	}

	word = strings.ToLower(word)

	currentNode := t.root

	for i := 0; i < len(word); {
		char := word[i]

		// if a child node exists for char
		if child, ok := currentNode.children[char]; ok {
			// if the edge label for the child node matches what's left of the word
			// edge label: 'abc', what's left of the word: 'abc'
			if child.edgeLabel == word[i:] {
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

			break
		}
	}

	return
}

// AddWords adds the given words to the trie t
func (t *trie) AddWords(words []string) {
	for i := 0; i < len(words); i++ {
		t.AddWord(words[i])
	}

	return
}

// DeleteWord removes the given word from the trie t
func (t *trie) DeleteWord(word string) error {
	// TODO implement
	return nil
}

// DeleteWords removes the given words from the trie t
func (t *trie) DeleteWords(words []string) error {
	for i := 0; i < len(words); i++ {
		if err := t.DeleteWord(words[i]); err != nil {
			return err
		}
	}

	return nil
}

// HasWord checks to see if the given word exists in the trie t
func (t trie) HasWord(word string) (bool, error) {
	// TODO implement
	return false, nil
}

// GetWordsFromPrefix returns all the words in trie t that have the given prefix
func (t trie) GetWordsFromPrefix(prefix string) ([]string, error) {
	// TODO implement
	words := make([]string, 0)
	return words, nil
}
