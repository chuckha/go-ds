# Data structures in Go

## Trie
[Trie](http://en.wikipedia.org/wiki/Trie)
Using the Trie:

```.go
        t := trie.NewTrie()
        t.Insert("tea")
        t.Insert("ten")
        t.Insert("tree")
        t.Find("tree") // is true
```
