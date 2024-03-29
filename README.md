# Go-Autocomplete

![Build](https://github.com/mattroseman/go-autocomplete/workflows/Go%20Build/badge.svg)
![Test](https://github.com/mattroseman/go-autocomplete/workflows/Go%20Test/badge.svg)

Autocomplete library built with Trie data structures.

## TESTS

`go test -v ./...` will run all tests.

### TEST COVERAGE

`go test ./... -coverprofile=coverage.out` will create the coverage report file.
`go tool cover -html coverage.out` will open up an HTML view of the coverage report in your browser

## BENCHMARKS

Before running benchmarks, unzip the `data/words.zip` file to `data/words.txt`. This file contains a list of english words used to benchmark with many words.

`go test -bench=. ./...` will run benchmarks.
- **BenchmarkAddWord** benchmarks how long it takes on average to add a word to the trie
- **BenchmarkAddWords**  benchmarks how long it takes to add the english dictionary to the trie
