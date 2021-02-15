package trie

type hasWordTest struct {
	name string
	trie Trie
	input string
	expected bool
}

var hasWordTestTrie = Trie{
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
