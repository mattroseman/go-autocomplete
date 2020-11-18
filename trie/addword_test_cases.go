package trie

type addWordTest struct {
	input string
	expected Trie
}

var addWordTestCases = []addWordTest{
	{
		"cucumber",
		Trie{
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
		Trie{
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
		Trie{
			root: &node{
				edgeLabel: "",
				endOfWord: false,
				children: make(map[byte]*node),
			},
		},
	},
	{
		"cUCumBeR",
		Trie{
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
	expected Trie
}

var addWordsTestCases = []addWordsTest{
	{
		[]string{},
		Trie{
			root: &node{
				edgeLabel: "",
				endOfWord: false,
				children: make(map[byte]*node),
			},
		},
	},
	{
		[]string{"cucumber", "banana", "bacon"},
		Trie{
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
		Trie{
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
		Trie{
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
		Trie{
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
