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
				children: make(map[byte]*node),
				endOfWord: false,
			},
		},
	},
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
	{
		[]string{"benchpress", "bench"},
		trie{
			root: &node{
				edgeLabel: "",
				children: map[byte]*node{
					'b': &node{
						edgeLabel: "bench",
						children: map[byte]*node{
							'p': &node{
								edgeLabel: "press",
								children: make(map[byte]*node),
								endOfWord: true,
							},
						},
						endOfWord: true,
					},
				},
				endOfWord: false,
			},
		},
	},
	{
		[]string{"banana", "banner", "ban"},
		trie{
			root: &node{
				edgeLabel: "",
				children: map[byte]*node{
					'b': &node{
						edgeLabel: "ban",
						children: map[byte]*node{
							'a': &node{
								edgeLabel: "ana",
								children: make(map[byte]*node),
								endOfWord: true,
							},
							'n': &node{
								edgeLabel: "ner",
								children: make(map[byte]*node),
								endOfWord: true,
							},
						},
						endOfWord: true,
					},
				},
				endOfWord: false,
			},
		},
	},
	{
		[]string{"ban", "banana"},
		trie{
			root: &node{
				edgeLabel: "",
				children: map[byte]*node{
					'b': &node{
						edgeLabel: "ban",
						children: map[byte]*node{
							'a': &node{
								edgeLabel: "ana",
								children: make(map[byte]*node),
								endOfWord: true,
							},
						},
						endOfWord: true,
					},
				},
				endOfWord: false,
			},
		},
	},
}
