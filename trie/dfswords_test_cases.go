package trie

type dfsWordsTest struct {
	name string
	startNode node
	prefix string
	expected []string
}

var dfsWordsTestTrie = Trie{
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

var dfsWordsTestCases = []dfsWordsTest{
	{
		"nominal case",
		*dfsWordsTestTrie.root.children['b'],
		"ban",
		[]string{"banana", "banner"},
	},
	{
		"leaf node given as start",
		*dfsWordsTestTrie.root.children['b'].children['a'],
		"banana",
		[]string{"banana"},
	},
	{
		"empty prefix given",
		*dfsWordsTestTrie.root,
		"",
		[]string{"banana", "banner", "cucumber"},
	},
	{
		"root node given as start with made up prefix",
		*dfsWordsTestTrie.root,
		"test",
		[]string{"testbanana", "testbanner", "testcucumber"},
	},
}
