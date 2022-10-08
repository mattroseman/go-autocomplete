package trie

type hasWordTest struct {
	name     string
	trie     Trie
	input    string
	expected bool
}

var hasWordTestTrie = Trie{
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

var hasWordTestCases = []hasWordTest{
	{
		"nominal case",
		hasWordTestTrie,
		"cucumber",
		true,
	},
	{
		"word not in trie",
		hasWordTestTrie,
		"apple",
		false,
	},
	{
		"empty string",
		hasWordTestTrie,
		"",
		false,
	},
	{
		"node in trie but not endOfWord",
		hasWordTestTrie,
		"ban",
		false,
	},
	{
		"empty trie",
		*New(),
		"banana",
		false,
	},
	{
		"word with variable case",
		hasWordTestTrie,
		"CuCuMBEr",
		true,
	},
}
