package trie

type Trie struct {
	End  bool
	Node map[string]Trie
}

func NewTrie() Trie {
	return Trie{false, make(map[string]Trie)}
}

func (t Trie) Insert(key string) {
	cn := t
	for i := 0; i < len(key); i++ {
		s := key[i : i+1]
		_, ok := cn.Node[s]
		if !ok {
			if i == len(key) - 1 {
				cn.Node[s] = Trie{true, make(map[string]Trie)}
			} else {
				cn.Node[s] = Trie{false, make(map[string]Trie)}
			}
		}
		cn = cn.Node[s]
	}
}

func (t Trie) Find(key string) bool {
	cn := t
	for i := 0; i < len(key); i++ {
		s := key[i : i+1]
		_, ok := cn.Node[s]
		if !ok {
			return false
		}
		cn = cn.Node[s]
	}
	return cn.End
}

