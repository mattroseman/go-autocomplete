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

func (t *Trie) deepCopy() *Trie {
	newTrie := New()
	newTrie.root = t.root.deepCopy()
	return newTrie
}

func (n *node) deepCopy() *node {
	newNode := &node{
		edgeLabel: n.edgeLabel,
		endOfWord: n.endOfWord,
		children: make(map[byte]*node),
	}
	for key, child := range n.children {
		newNode.children[key] = child.deepCopy()
	}

	return newNode
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

// traverseTrie traverses down the trie with the given word.
// It returns the final node found or nil if the given word doesn't exist in the trie (even as a prefix).
// It also returns a string representing what's left of the given word not covered up till the returned node.
// Ex 1: t.traverseTrie("foobar") => (nodeA, "") means trie t contains the word "foobar" up till nodeA with nothing left over
//       (if foobar is an actual word can be determined by checking nodeA.endOfWord)
// Ex 2: t.traverseTrie("foobar") => (nodeB, "bar") means trie t contains the word "foo" up till nodeB with "bar" left over
//		 ("foobar" could still exist in trie t as a prefix, just not a word,
//		  "foobarxyz" could exist as a word but the edgle label "barxyz" differs from what's left of the word "bar")
// Ex 3: t.traverseTrie("foobar") => (nil, "") means trie t does not contain "foobar" as a word or even a prefix.
func (t *Trie) traverseTrie(word string) (*node, string) {
	word = strings.ToLower(word)
	currentNode := t.root

	// a blank word ends at the root node of the trie
	if word == "" {
		return currentNode, ""
	}

	for i := 0; i < len(word); {
		char := word[i]

		if child, ok := currentNode.children[char]; ok {
			// if the edge label matches what's left of the word, than the child is the final node
			if (child.edgeLabel == word[i:]) {
				// the entire word is in the trie and ends at the child node with nothing leftover
				return child, ""
			}

			commonPrefix := getCommonPrefix(child.edgeLabel, word[i:])

			// if the edge label for the child contains the entirety of what's left of the word plus some extra
			// edge label: 'abcdef', what's left of the word: 'abc'
			if len(commonPrefix) < len(child.edgeLabel) && len(commonPrefix) == len(word[i:]) {
				// the word could exist in the trie as a prefix, but it does not end on a node, but partway through an edge label
				return currentNode, word[i:]
			}

			// if the edge label and what's left of the word share a common prefix, but differ at some point
			// edge label: 'abcdef', what's left of the word: 'abcxyz'
			if len(commonPrefix) < len(child.edgeLabel) && len(commonPrefix) < len(word[i:]) {
				// the word doesn't exist in the trie as a word or prefix
				return nil, ""
			}

			// the last option is that what's left of the word contains the entirety of the edge label plus some extra
			// edge label: 'abc', what's left of the word: 'abcdef'
			// follow the edge to the next node and increment i to account for the edge's label
			i += len(child.edgeLabel)
			currentNode = child
		} else {
			break
		}
	}

	// there are no words in t that begin with the same character word begins with
	return nil, ""
}

// HasWord checks to see if the given word exists in the trie t.
// It returns a boolean of true if the given word exists in the trie, false otherwise.
func (t Trie) HasWord(word string) bool {
	endNode, leftover := t.traverseTrie(word)

	// if the given word ends cleanly on a node and that node is an endOfWord node the word exists in the trie t
	return endNode != nil && endNode.endOfWord && leftover == ""
}

// DeleteWord removes the given word from the trie t.
// It returns a boolean of true if the given word was found and removed, false if the given word couldn't be found.
func (t *Trie) DeleteWord(word string) bool {
	if len(word) == 0 {
		return false
	}

	wordNode, leftover := t.traverseTrie(word)

	// if the given word isn't a node in the trie or isn't an endOfWord node
	if wordNode == nil || len(leftover) > 0 || !wordNode.endOfWord {
		return false
	}

	parentNode, _ := t.traverseTrie(word[:len(word) - 1])

	// if node doesn't have any children, delete it from the parent node
	if len(wordNode.children) == 0 {
		delete(parentNode.children, wordNode.edgeLabel[0])
		return true
	}

	// if the node has one child merge the child's edge label with the parent edge label
	if len(wordNode.children) == 1 {
		for key := range wordNode.children {
			childNode := wordNode.children[key]
			parentNode.children[wordNode.edgeLabel[0]] = childNode
			childNode.edgeLabel = wordNode.edgeLabel + childNode.edgeLabel
		}

		return true
	}

	// if the node has multiple children mark the node as not endOfWord
	wordNode.endOfWord = false

	return true
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

	prefixEndNode, leftover := t.traverseTrie(prefix)
	// if the prefix isn't even found in the trie, return the empty array
	if prefixEndNode == nil {
		return words
	}

	if len(leftover) > 0 {
		nextNode := prefixEndNode.children[leftover[0]]
		newPrefix := prefix + nextNode.edgeLabel[len(leftover):]

		return nextNode.dfsWords(newPrefix)
	}

	return prefixEndNode.dfsWords(prefix)
}
