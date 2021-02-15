# Go-Autocomplete

![Build](https://github.com/mattroseman/go-autocomplete/workflows/Go%20Build/badge.svg)
![Test](https://github.com/mattroseman/go-autocomplete/workflows/Go%20Test/badge.svg)

Autocomplete library built with Trie data structures.

## TESTS

`go test -v ./...` will run all tests.

`go test -bench=. ./...` will run benchmarks.
**BenchmarkAddWord** benchmarks how long it takes on average to add a word to the trie
**BenchmarkAddWords**  benchmarks how long it takes to add the english dictionary to the trie
