package trie

type deleteWordTest struct {
	name     string
	trie     Trie
	input    string
	succeeds bool
	expected Trie
}

var deleteWordTestTrie = Trie{
	root: &node{
		children: map[byte]*node{
			'c': {
				edgeLabel: "cucumber",
				endOfWord: true,
				children: map[byte]*node{
					's': {
						edgeLabel: "s",
						endOfWord: true,
						children:  make(map[byte]*node),
					},
				},
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
			'r': {
				edgeLabel: "real",
				endOfWord: true,
				children: map[byte]*node{
					'y': {
						edgeLabel: "y",
						endOfWord: true,
						children:  make(map[byte]*node),
					},
					'i': {
						edgeLabel: "ity",
						endOfWord: true,
						children:  make(map[byte]*node),
					},
				},
			},
		},
	},
}

var deleteWordTestTrie2 = Trie{
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
			'r': {
				edgeLabel: "real",
				endOfWord: true,
				children: map[byte]*node{
					'y': {
						edgeLabel: "y",
						endOfWord: true,
						children:  make(map[byte]*node),
					},
					'i': {
						edgeLabel: "ity",
						endOfWord: true,
						children:  make(map[byte]*node),
					},
				},
			},
		},
	},
}

var deleteWordTestTrie3 = Trie{
	root: &node{
		children: map[byte]*node{
			'c': {
				edgeLabel: "cucumbers",
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
			'r': {
				edgeLabel: "real",
				endOfWord: true,
				children: map[byte]*node{
					'y': {
						edgeLabel: "y",
						endOfWord: true,
						children:  make(map[byte]*node),
					},
					'i': {
						edgeLabel: "ity",
						endOfWord: true,
						children:  make(map[byte]*node),
					},
				},
			},
		},
	},
}

var deleteWordTestTrie4 = Trie{
	root: &node{
		children: map[byte]*node{
			'c': {
				edgeLabel: "cucumber",
				endOfWord: true,
				children: map[byte]*node{
					's': {
						edgeLabel: "s",
						endOfWord: true,
						children:  make(map[byte]*node),
					},
				},
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
			'r': {
				edgeLabel: "real",
				endOfWord: false,
				children: map[byte]*node{
					'y': {
						edgeLabel: "y",
						endOfWord: true,
						children:  make(map[byte]*node),
					},
					'i': {
						edgeLabel: "ity",
						endOfWord: true,
						children:  make(map[byte]*node),
					},
				},
			},
		},
	},
}

var deleteWordTestCases = []deleteWordTest{
	{
		"nominal case",
		deleteWordTestTrie,
		"cucumbers",
		true,
		deleteWordTestTrie2,
	},
	{
		"word not found",
		deleteWordTestTrie,
		"foobar",
		false,
		deleteWordTestTrie,
	},
	{
		"empty word",
		deleteWordTestTrie,
		"",
		false,
		deleteWordTestTrie,
	},
	{
		"word ends partway through edge",
		deleteWordTestTrie,
		"banan",
		false,
		deleteWordTestTrie,
	},
	{
		"word ends on node that isn't endOfWord",
		deleteWordTestTrie,
		"ban",
		false,
		deleteWordTestTrie,
	},
	{
		"word has one child",
		deleteWordTestTrie,
		"cucumber",
		true,
		deleteWordTestTrie3,
	},
	{
		"word has multiple children",
		deleteWordTestTrie,
		"real",
		true,
		deleteWordTestTrie4,
	},
}
