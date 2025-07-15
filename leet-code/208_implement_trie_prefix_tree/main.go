package _08_implement_trie_prefix_tree

// need code の解説見てそのままやった
// すごいなTrie
type Trie struct {
	children  map[rune]*Trie
	endOfWord bool
}

func Constructor() Trie {
	return Trie{
		children:  make(map[rune]*Trie),
		endOfWord: false,
	}
}

func (t *Trie) Insert(word string) {
	cur := t

	for _, c := range word {
		if _, ok := cur.children[c]; !ok {
			// haven't existed yet. add c as a node
			cur.children[c] = &Trie{children: make(map[rune]*Trie)}
		}
		cur = cur.children[c]
	}
	cur.endOfWord = true
}

func (t *Trie) Search(word string) bool {
	cur := t
	for _, c := range word {
		if _, ok := cur.children[c]; !ok {
			return false
		}
		cur = cur.children[c]
	}
	return cur.endOfWord
}

func (t *Trie) StartsWith(prefix string) bool {
	cur := t
	for _, c := range prefix {
		if _, ok := cur.children[c]; !ok {
			return false
		}
		cur = cur.children[c]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
