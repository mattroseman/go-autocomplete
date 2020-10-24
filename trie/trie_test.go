package trie

import "testing"

func TestAddWord(t *testing.T) {
	trie := NewTrie()
	testWord := "cucumber"

	err := trie.AddWord(testWord)
	if err != nil {
		t.Errorf("Something went wrong adding word 'cucumber' to trie\n%s", err)
	}

	currentNode := trie.root

	for i := 0; i < len(testWord); {
		if edge, ok := currentNode.children[testWord[i]]; ok {
			if len(testWord[i:]) <= len(edge.label) && testWord[i:i + len(edge.label)] == edge.label {
				currentNode = edge.target
				i += len(edge.label)
			} else {
				t.Errorf("The added word, %s, was not found in the trie after calling AddWord", testWord)
				break
			}
		} else {
			t.Errorf("The added word, %s, was not found in the trie after calling AddWord", testWord)
			break
		}
	}
}
