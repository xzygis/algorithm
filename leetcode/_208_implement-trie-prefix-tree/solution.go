package _208_implement_trie_prefix_tree

type Trie struct {
	children [26]*Trie //长度为26的数组，方便后续写入数据
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	cur := t
	for _, ch := range word {
		index := ch - 'a'
		if cur.children[index] == nil {
			cur.children[index] = &Trie{}
		}
		cur = cur.children[index]
	}
	cur.isEnd = true
}

func (t *Trie) Search(word string) bool {
	if len(word) == 0 {
		return false
	}

	node := t.searchPrefix(word)
	return node != nil && node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.searchPrefix(prefix) != nil
}

func (t *Trie) searchPrefix(prefix string) *Trie {
	cur := t
	for _, ch := range prefix {
		index := ch - 'a'
		if cur.children[index] == nil {
			return nil
		}
		cur = cur.children[index]
	}
	return cur
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
