package trie

type getWordsFromPrefixTest struct {
	name string
	trie *Trie
	prefix string
	expected []string
}

var getWordsFromPrefixTestTrie = Trie{
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

var getWordsFromPrefixTestCases = []getWordsFromPrefixTest{
	{
		"nominal case",
		&getWordsFromPrefixTestTrie,
		"ban",
		[]string{"banana", "banner"},
	},
	{
		"empty prefix",
		&getWordsFromPrefixTestTrie,
		"",
		[]string{"banana", "banner", "cucumber"},
	},
	{
		"not found prefix",
		&getWordsFromPrefixTestTrie,
		"foobar",
		[]string{},
	},
	{
		"prefix doesn't end on node",
		&getWordsFromPrefixTestTrie,
		"banan",
		[]string{"banana"},
	},
}
