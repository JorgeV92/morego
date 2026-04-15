# morego

A small Go project for practicing algorithms 


## Run a problem

```bash
go run ./cmd/arrays -problem three-sum
go run ./cmd/arrays -problem array-change

go run ./cmd/graph -problem topological-sort
go run ./cmd/graph -problem obstacle-removal
go run ./cmd/graph -problem surrounding-xo

go run ./cmd/data -problem treap
```

Or with `make`:

```bash
make run-arrays ARRAY_PROBLEM=three-sum
make run-graph GRAPH_PROBLEM=surrounding-xo
make run-data DATA_PROBLEM=treap
```

## Parse a Go AST

```bash
go run ./cmd/parsing -file algorithms/parsing/ast.go
go run ./cmd/parsing -file algorithms/data/mst.go
```

Or with `make`:

```bash
make run-parsing PARSING_FILE=algorithms/data/mst.go
```

## Run tests

```bash
go test ./...
make test
```

## Add a new algorithm

1. Add the pure function to `algorithms/graph/problems.go`.
2. Add a test in `algorithms/graph/problems_test.go`.
3. Register a sample case in `cmd/algorithms/main.go`.

## Treap example

```bash
go run ./cmd/data -problem treap
```

Sample output:

```text
after insert: [1 2 3 4 5]
size: 5
sum: 15
after erase 3: [1 2 4 5]
left split: [1 2]
right split: [4 5]
merged again: [1 2 4 5]
```
