package trie

type Trie struct {
	Value string
	Node  map[string]Trie
}

func NewTrie() Trie {
	return Trie{"", make(map[string]Trie)}
}

func (t Trie) Insert(key string) {
	cn := t
	for i := 0; i < len(key); i++ {
		s := key[i:i+1]
		_, ok := cn.Node[s]
		if !ok {
			cn.Node[s] = Trie{key[:i+1], make(map[string]Trie)}
		}
		cn = cn.Node[s]
	}
}

func (t Trie) Find(key string) bool {
	cn := t
	for i := 0; i < len(key); i++ {
		s := key[i:i+1]
		_, ok := cn.Node[s]
		if !ok {
			return false
		}
		cn = cn.Node[s]
	}
	return true
}

