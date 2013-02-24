# Data structures in Go
Installing this package:

    go get github.com/ggnextmap/go-ds/...

## [Trie](http://en.wikipedia.org/wiki/Trie)
### Importing the Trie:

    import "github.com/ggnextmap/go-ds/trie"

### Using the Trie:

```.go
t := trie.NewTrie()
t.Insert("tea")
t.Insert("ten")
t.Insert("tree")
t.Find("tree") // is true
```

### Testing and Benchmarking:

    go test github.com/ggnextmap/go-ds/trie
    go test -test.bench=".*" github.com/ggnextmap/go-ds/trie

### ToDo

* Removing a key
* Passing tests
