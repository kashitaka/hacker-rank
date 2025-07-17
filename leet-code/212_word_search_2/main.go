package _12_word_search_2

// むずい。解けるわけねーや
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
	}
}

func (r *TrieNode) Add(word string) {
	cur := r
	for _, c := range word {
		if cur.children[c] == nil {
			cur.children[c] = &TrieNode{
				children: make(map[rune]*TrieNode),
			}
		}
		cur = cur.children[c]
	}
	cur.isEnd = true
}

func findWords(board [][]byte, words []string) []string {
	ROWS, COLS := len(board), len(board[0])
	trie := NewTrieNode()
	// O(Sum(len of each word))
	for _, word := range words {
		trie.Add(word)
	}

	visited := make(map[[2]int]bool)
	wordSet := make(map[string]struct{})

	var dfs func(row int, col int, cur *TrieNode, runes []rune)
	dfs = func(row int, col int, cur *TrieNode, runes []rune) {
		if row < 0 || col < 0 || row == ROWS || col == COLS ||
			visited[[2]int{row, col}] || cur.children[rune(board[row][col])] == nil {
			return
		}
		// word is valid so far
		runes = append(runes, rune(board[row][col]))
		cur = cur.children[rune(board[row][col])]
		visited[[2]int{row, col}] = true
		if cur.isEnd {
			wordSet[string(runes)] = struct{}{}
		}

		dfs(row-1, col, cur, runes)
		dfs(row+1, col, cur, runes)
		dfs(row, col-1, cur, runes)
		dfs(row, col+1, cur, runes)

		// backtracking
		visited[[2]int{row, col}] = false
	}

	for r, row := range board {
		for c, _ := range row {
			dfs(r, c, trie, []rune{})
		}
	}

	res := []string{}
	for k, _ := range wordSet {
		res = append(res, k)
	}
	return res
}
