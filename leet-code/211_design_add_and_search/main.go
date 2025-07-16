package _11_design_add_and_search

// いやもうこれむず過ぎるだろ!!!!! いい加減にしろ！！
// Neet Code の解法見た上でそっからさらに1時間かかったわ！

type WordDictionary struct {
	hash  map[rune]*WordDictionary
	isEnd bool
}

func Constructor() WordDictionary {
	return WordDictionary{
		hash: make(map[rune]*WordDictionary),
	}
}

func (this *WordDictionary) AddWord(word string) {
	cur := this
	for _, c := range word {
		if _, ok := cur.hash[c]; !ok {
			cur.hash[c] = &WordDictionary{
				hash: make(map[rune]*WordDictionary),
			}
		}
		cur = cur.hash[c]
	}
	cur.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	var dfs func(j int, root *WordDictionary) bool
	dfs = func(j int, root *WordDictionary) bool {
		cur := root
		for i, c := range word[j:] {
			if c == '.' {
				for _, child := range cur.hash {
					if dfs(i+j+1, child) {
						return true
					}
				}
				return false
			} else {
				if _, ok := cur.hash[c]; !ok {
					return false
				}
				cur = cur.hash[c]
			}
		}
		return cur.isEnd
	}
	return dfs(0, this)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
