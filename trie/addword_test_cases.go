package trie

var addWordTestCases = []struct{
	input string
	expected trie
}{
	{
		"cucumber",
		trie{
			root: &node{
				edgeLabel: "",
				children: map[byte]*node{
					'c': &node{
						edgeLabel: "cucumber",
						children: make(map[byte]*node),
						endOfWord: true,
					},
				},
				endOfWord: false,
			},
		},
	},
	{
		"test",
		trie{
			root: &node{
				children: map[byte]*node{
					't': &node{
						edgeLabel: "test",
						children: make(map[byte]*node),
						endOfWord: true,
					},
				},
			},
		},
	},
	{
		"",
		trie{
			root: &node{
				children: make(map[byte]*node),
				endOfWord: false,
			},
		},
	},
	{
		"cUCumBeR",
		trie{
			root: &node{
				children: map[byte]*node{
					'c': &node{
						edgeLabel: "cucumber",
						children: make(map[byte]*node),
						endOfWord: true,
					},
				},
			},
		},
	},
}

var addWordsTestCases = []struct{
	input []string
	expected trie
}{
	{
		[]string{"cucumber", "banana", "bacon"},
		trie{
			root: &node{
				edgeLabel: "",
				children: map[byte]*node{
					'c': &node{
						edgeLabel: "cucumber",
						children: make(map[byte]*node),
						endOfWord: true,
					},
					'b': &node{
						edgeLabel: "ba",
						children: map[byte]*node{
							'n': &node{
								edgeLabel: "nana",
								children: make(map[byte]*node),
								endOfWord: true,
							},
							'c': &node{
								edgeLabel: "con",
								children: make(map[byte]*node),
								endOfWord: true,
							},
						},
						endOfWord: false,
					},
				},
				endOfWord: false,
			},
		},
	},
}
