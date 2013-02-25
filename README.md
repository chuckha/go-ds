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

## [Suffix Tree](http://en.wikipedia.org/wiki/Suffix_tree)
### Importing the Suffix Tree:

    import "github.com/ggnextmap/go-ds/suffixtree"

## Using the Suffix Tree:

```.go
```

## Testing and Benchmarking:

    go test github.com/ggnextmap/go-ds/suffixtree
    go test -test.bench=".*" github.com/ggnextmap/go-ds/suffixtree

## ToDo
