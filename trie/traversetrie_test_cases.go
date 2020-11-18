package trie

type traverseTrieTest struct {
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
		traverseTrieTestTrie,
		"banana",
		true,
		traverseTrieTestTrie.root.children['b'].children['a'],
	},
	{
		traverseTrieTestTrie,
		"superman",
		false,
		nil,
	},
	{
		traverseTrieTestTrie,
		"ban",
		false,
		nil,
	},
	{
		traverseTrieTestTrie,
		"bananas",
		false,
		nil,
	},
	{
		traverseTrieTestTrie,
		"BaNanA",
		true,
		traverseTrieTestTrie.root.children['b'].children['a'],
	},
	{
		traverseTrieTestTrie,
		"",
		false,
		nil,
	},
	{
		*New(),
		"banana",
		false,
		nil,
	},
}
