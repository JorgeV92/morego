# MoreGo

A small Go project for practicing algorithms with:

- pure functions inside packages
- separate CLI runners by topic
- table-friendly unit tests

## Project structure

```text
.
├── algorithms
│   ├── arrays
│   │   ├── problems.go
│   │   └── problems_test.go
│   └── graph
│       ├── problems.go
│       └── problems_test.go
├── cmd
│   ├── arrays
│   │   └── main.go
│   └── graph
│       └── main.go
├── go.mod
└── Makefile
```

## Run array problems

```bash
go run ./cmd/arrays -problem three-sum
go run ./cmd/arrays -problem array-change
```

Or with `make`:

```bash
make run-arrays ARRAY_PROBLEM=three-sum
```

## Run graph problems

```bash
go run ./cmd/graph -problem obstacle-removal
go run ./cmd/graph -problem topological-sort
go run ./cmd/graph -problem surrounding-xo
```

Or with `make`:

```bash
make run-graph GRAPH_PROBLEM=topological-sort
```

## Run tests

```bash
go test ./...
make test
```

## Add a new algorithm

1. Pick the topic folder such as `algorithms/graph` or `algorithms/arrays`.
2. Add the pure function to that package's `problems.go`.
3. Add a test in that package's `problems_test.go`.
4. Register a sample case in the matching runner, like `cmd/graph/main.go` or `cmd/arrays/main.go`.

That keeps your logic, test cases, and demo runner separated cleanly.
