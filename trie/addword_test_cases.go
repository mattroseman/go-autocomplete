package trie

type addWordTest struct {
	name     string
	input    string
	expected Trie
}

var addWordTestCases = []addWordTest{
	{
		"nominal case",
		"cucumber",
		Trie{
			root: &node{
				children: map[byte]*node{
					'c': {
						edgeLabel: "cucumber",
						endOfWord: true,
						children:  make(map[byte]*node),
					},
				},
			},
		},
	},
	{
		"blank word",
		"",
		Trie{
			root: &node{
				edgeLabel: "",
				endOfWord: false,
				children:  make(map[byte]*node),
			},
		},
	},
	{
		"variable case word",
		"cUCumBeR",
		Trie{
			root: &node{
				children: map[byte]*node{
					'c': {
						edgeLabel: "cucumber",
						endOfWord: true,
						children:  make(map[byte]*node),
					},
				},
			},
		},
	},
}

type addWordsTest struct {
	name     string
	input    []string
	expected Trie
}

var addWordsTestCases = []addWordsTest{
	{
		"empty array",
		[]string{},
		Trie{
			root: &node{
				edgeLabel: "",
				endOfWord: false,
				children:  make(map[byte]*node),
			},
		},
	},
	{
		"words with shared prefix",
		[]string{"cucumber", "banana", "bacon"},
		Trie{
			root: &node{
				children: map[byte]*node{
					'c': {
						edgeLabel: "cucumber",
						endOfWord: true,
						children:  make(map[byte]*node),
					},
					'b': {
						edgeLabel: "ba",
						endOfWord: false,
						children: map[byte]*node{
							'n': {
								edgeLabel: "nana",
								endOfWord: true,
								children:  make(map[byte]*node),
							},
							'c': {
								edgeLabel: "con",
								endOfWord: true,
								children:  make(map[byte]*node),
							},
						},
					},
				},
			},
		},
	},
	{
		"word with node already existing, but not endOfWord",
		[]string{"benchpress", "bench"},
		Trie{
			root: &node{
				children: map[byte]*node{
					'b': {
						edgeLabel: "bench",
						endOfWord: true,
						children: map[byte]*node{
							'p': {
								edgeLabel: "press",
								endOfWord: true,
								children:  make(map[byte]*node),
							},
						},
					},
				},
			},
		},
	},
	{
		"shared prefix and already existing node, but not endOfWord",
		[]string{"banana", "banner", "ban"},
		Trie{
			root: &node{
				children: map[byte]*node{
					'b': {
						edgeLabel: "ban",
						endOfWord: true,
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
		},
	},
	{
		"new node off existing node",
		[]string{"ban", "banana"},
		Trie{
			root: &node{
				children: map[byte]*node{
					'b': {
						edgeLabel: "ban",
						endOfWord: true,
						children: map[byte]*node{
							'a': {
								edgeLabel: "ana",
								endOfWord: true,
								children:  make(map[byte]*node),
							},
						},
					},
				},
			},
		},
	},
}
