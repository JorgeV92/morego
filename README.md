# MoreGo

A small Go project for practicing algorithms 


## Run a problem

```bash
go run ./cmd/algorithms -problem obstacle-removal
go run ./cmd/algorithms -problem topological-sort
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

