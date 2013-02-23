package trie

import (
	"testing"
)

func TestT(t *testing.T) {
	trie := NewTrie()
	vals := []string{"A", "to", "tea", "ted", "ten", "i", "in", "inn"}

	trie.Insert(vals[0])

	if trie.Find("A") != true {
		t.Errorf("Should have found 'A'")
	}

	trie.Insert(vals[2])

	if trie.Find("tea") != true {
		t.Errorf("Should have found 'tea'")
	}

	if trie.Find("nope") == true {
		t.Errorf("Should not have but did find 'nope'")
	}
}
