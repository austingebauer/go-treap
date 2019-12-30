# go-treap [![GoReportCard](https://goreportcard.com/badge/github.com/austingebauer/go-treap)](https://goreportcard.com/report/github.com/austingebauer/go-treap) [![GoDoc](https://godoc.org/github.com/austingebauer/go-treap?status.svg)](https://godoc.org/github.com/austingebauer/go-treap)

A Go library that provides a balanced binary search tree 
by using binary heap properties and randomization.

The height of the treap is proportional to the logarithm of the number 
of values inserted with high probability. This characteristic means that, 
on average, each `Search()`, `Insert()`, or `Delete()` operation takes logarithmic 
time to perform.  

## Installation

To install `go-treap`, use `go get`.

```bash
go get github.com/austingebauer/go-treap
```

Then import the library into your Go program.

```go
import "github.com/austingebauer/go-treap"
```

## Usage

### API

The `go-treap` API provides `Search()`, `Insert()`, and `Delete()` functions with 
`O(log n)` time complexity on average.  

Please refer to the [GoDoc](https://godoc.org/github.com/austingebauer/go-treap) for 
additional API documentation of the library.

**Example 1**: Basic usage
```go
trp := NewTreap()

trp.Insert("c")
trp.Insert("b")
trp.Insert("a")

trp.Search("a") // true
trp.Search("d") // false

trp.Delete("b")
trp.Delete("c")
trp.Delete("a")
```

### Behavior

For a simple explanation of the treap data structure, I recommend reading 
[Julia Evan's Blog Post on Treaps](https://jvns.ca/blog/2017/09/09/data-structure--the-treap-/).

## Contributing

Pull requests are welcome. 

For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests along with changes.

## License

[MIT](LICENSE)
