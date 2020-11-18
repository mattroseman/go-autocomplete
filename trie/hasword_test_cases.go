package trie

type hasWordTest struct {
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
		hasWordTestTrie,
		"cucumber",
		true,
	},
	{
		hasWordTestTrie,
		"apple",
		false,
	},
	{
		hasWordTestTrie,
		"",
		false,
	},
	{
		hasWordTestTrie,
		"ban",
		false,
	},
	{
		*New(),
		"banana",
		false,
	},
	{
		hasWordTestTrie,
		"CuCuMBEr",
		true,
	},
}
