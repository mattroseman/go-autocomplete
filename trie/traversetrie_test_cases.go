package trie

type traverseTrieTest struct {
	name             string
	trie             Trie
	input            string
	expectedEndNode  *node
	expectedLeftover string
}

var traverseTrieTestTrie = Trie{
	root: &node{
		children: map[byte]*node{
			'c': {
				edgeLabel: "cucumber",
				endOfWord: true,
				children:  make(map[byte]*node),
			},
			'b': {
				edgeLabel: "ban",
				endOfWord: false,
				children: map[byte]*node{
					'a': {
						edgeLabel: "ana",
						endOfWord: true,
						children:  make(map[byte]*node),
					},
					'n': {
						edgeLabel: "ner",
						endOfWord: true,
						children:  make(map[byte]*node),
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
		traverseTrieTestTrie.root.children['b'].children['a'],
		"",
	},
	{
		"word not in trie",
		traverseTrieTestTrie,
		"superman",
		nil,
		"",
	},
	{
		"node in trie but not endOfWord",
		traverseTrieTestTrie,
		"ban",
		traverseTrieTestTrie.root.children['b'],
		"",
	},
	{
		"word has extra character",
		traverseTrieTestTrie,
		"bananas",
		nil,
		"",
	},
	{
		"encounters edgeLabel that matches first character but differs later",
		traverseTrieTestTrie,
		"banal",
		nil,
		"",
	},
	{
		"word but with variable case",
		traverseTrieTestTrie,
		"BaNanA",
		traverseTrieTestTrie.root.children['b'].children['a'],
		"",
	},
	{
		"empty string",
		traverseTrieTestTrie,
		"",
		traverseTrieTestTrie.root,
		"",
	},
	{
		"empty trie",
		*New(),
		"banana",
		nil,
		"",
	},
	{
		"word ends on edgelabel",
		traverseTrieTestTrie,
		"banan",
		traverseTrieTestTrie.root.children['b'],
		"an",
	},
}
