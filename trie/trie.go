// Package trie implements a minified Deterministic Acyclic Finite State Automata (DAFSA).
// https://en.wikipedia.org/wiki/Deterministic_acyclic_finite_state_automaton
// A DAFSA data structure is useful for getting possible given a certain prefix which in turn can be used for autocomplete.
package trie

import "strings"

type edge struct {
	label string
	target node
}

type node struct {
	children map[byte]edge // maps the Edge's label's first character to the Edge itself
	endOfWord bool
}

type trie struct {
	root node
}

func NewTrie() trie {
	return trie{
		root: node{
			children: make(map[byte]edge),
			endOfWord: false,
		},
	}
}

// AddWord adds a given word to the trie t
func (t *trie) AddWord(word string) error {
	word = strings.ToLower(word)

	t.root.children[word[0]] = edge{
		label: word,
		target: node{
			endOfWord: true,
		},
	}

	return nil
}

// AddWords adds the given words to the trie t
func (t *trie) AddWords(words []string) error {
	for i := 0; i < len(words); i++ {
		if err := t.AddWord(words[i]); err != nil {
			return err
		}
	}

	return nil
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
