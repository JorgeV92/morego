# morego

A small Go project for practicing algorithms 


## Run a problem

```bash
go run ./cmd/arrays -problem three-sum
go run ./cmd/arrays -problem array-change

go run ./cmd/graph -problem topological-sort
go run ./cmd/graph -problem obstacle-removal
go run ./cmd/graph -problem surrounding-xo
```

Or with `make`:

```bash
make run PROBLEM=surrounding-xo
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

