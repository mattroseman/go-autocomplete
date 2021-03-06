package trie

type traverseTrieTest struct {
	name string
	trie Trie
	input string
	succeeds bool
	expected *node
}

var traverseTrieTestTrie = Trie{
	root: &node{
		children: map[byte]*node{
			'c': &node{
				edgeLabel: "cucumber",
				endOfWord: true,
				children: make(map[byte]*node),
			},
			'b': &node{
				edgeLabel: "ban",
				endOfWord: false,
				children: map[byte]*node{
					'a': &node{
						edgeLabel: "ana",
						endOfWord: true,
						children: make(map[byte]*node),
					},
					'n': &node{
						edgeLabel: "ner",
						endOfWord: true,
						children: make(map[byte]*node),
					},
				},
			},
		},
	},
}

var traverseTrieTestCases = []traverseTrieTest{
	{
		"nominal case",
		traverseTrieTestTrie,
		"banana",
		true,
		traverseTrieTestTrie.root.children['b'].children['a'],
	},
	{
		"word not in trie",
		traverseTrieTestTrie,
		"superman",
		false,
		nil,
	},
	{
		"node in trie but not endOfWord",
		traverseTrieTestTrie,
		"ban",
		false,
		nil,
	},
	{
		"word has extra character",
		traverseTrieTestTrie,
		"bananas",
		false,
		nil,
	},
	{
		"encounters edgeLabel that matches first character but differs later",
		traverseTrieTestTrie,
		"banal",
		false,
		nil,
	},
	{
		"word but with variable case",
		traverseTrieTestTrie,
		"BaNanA",
		true,
		traverseTrieTestTrie.root.children['b'].children['a'],
	},
	{
		"empty string",
		traverseTrieTestTrie,
		"",
		false,
		nil,
	},
	{
		"empty trie",
		*New(),
		"banana",
		false,
		nil,
	},
}
