package trie

type addWordTest struct {
	input string
	expected trie
}

var addWordTestCases = []addWordTest{
	{
		"cucumber",
		trie{
			root: &node{
				children: map[byte]*node{
					'c': &node{
						edgeLabel: "cucumber",
						endOfWord: true,
						children: make(map[byte]*node),
					},
				},
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
						endOfWord: true,
						children: make(map[byte]*node),
					},
				},
			},
		},
	},
	{
		"",
		trie{
			root: &node{
				edgeLabel: "",
				endOfWord: false,
				children: make(map[byte]*node),
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
						endOfWord: true,
						children: make(map[byte]*node),
					},
				},
			},
		},
	},
}

type addWordsTest struct {
	input []string
	expected trie
}

var addWordsTestCases = []addWordsTest{
	{
		[]string{},
		trie{
			root: &node{
				edgeLabel: "",
				endOfWord: false,
				children: make(map[byte]*node),
			},
		},
	},
	{
		[]string{"cucumber", "banana", "bacon"},
		trie{
			root: &node{
				children: map[byte]*node{
					'c': &node{
						edgeLabel: "cucumber",
						endOfWord: true,
						children: make(map[byte]*node),
					},
					'b': &node{
						edgeLabel: "ba",
						endOfWord: false,
						children: map[byte]*node{
							'n': &node{
								edgeLabel: "nana",
								endOfWord: true,
								children: make(map[byte]*node),
							},
							'c': &node{
								edgeLabel: "con",
								endOfWord: true,
								children: make(map[byte]*node),
							},
						},
					},
				},
			},
		},
	},
	{
		[]string{"benchpress", "bench"},
		trie{
			root: &node{
				children: map[byte]*node{
					'b': &node{
						edgeLabel: "bench",
						endOfWord: true,
						children: map[byte]*node{
							'p': &node{
								edgeLabel: "press",
								endOfWord: true,
								children: make(map[byte]*node),
							},
						},
					},
				},
			},
		},
	},
	{
		[]string{"banana", "banner", "ban"},
		trie{
			root: &node{
				children: map[byte]*node{
					'b': &node{
						edgeLabel: "ban",
						endOfWord: true,
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
		},
	},
	{
		[]string{"ban", "banana"},
		trie{
			root: &node{
				children: map[byte]*node{
					'b': &node{
						edgeLabel: "ban",
						endOfWord: true,
						children: map[byte]*node{
							'a': &node{
								edgeLabel: "ana",
								endOfWord: true,
								children: make(map[byte]*node),
							},
						},
					},
				},
			},
		},
	},
}
