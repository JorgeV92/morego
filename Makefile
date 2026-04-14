ARRAY_PROBLEM ?= array-change
GRAPH_PROBLEM ?= obstacle-removal
PARSING_FILE ?= algorithms/parsing/ast.go

.PHONY: run-arrays run-graph run-parsing test

run-arrays:
	go run ./cmd/arrays -problem $(ARRAY_PROBLEM)

run-graph:
	go run ./cmd/graph -problem $(GRAPH_PROBLEM)

run-parsing:
	go run ./cmd/parsing -file $(PARSING_FILE)

test:
	go test ./...
