package trie

import (
	"testing"
	"math/rand"
	"time"
)

func TestInsert(t *testing.T) {
	trie := NewTrie()

	trie.Insert("Ted")

	if len(trie.Node) != 1 {
		t.Errorf("Should have one node")
	}

	trie.Insert("Ten")
	if len(trie.Node) != 1 {
		t.Errorf("Should still have one node")
	}
}

func TestFind(t *testing.T) {
	trie := NewTrie()

	trie.Insert("Bill")

	if trie.Find("Bill") != true {
		t.Errorf("Should have found Bill")
	}

	if trie.Find("William") == true {
		t.Errorf("Should not have found William")
	}
}

func BenchmarkManyWordsTrie(b *testing.B) {
	trie := NewTrie()
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < b.N; i++ {
		word := ""
		for j := 0; j < (rand.Int() % 7) + 8; j++ {
			word += string((rand.Int() % 57) + 65)
		}
		trie.Insert(word)
	}
}
