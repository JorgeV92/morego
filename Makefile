ARRAY_PROBLEM ?= array-change
GRAPH_PROBLEM ?= obstacle-removal

.PHONY: run-arrays run-graph test

run-arrays:
	go run ./cmd/arrays -problem $(ARRAY_PROBLEM)

run-graph:
	go run ./cmd/graph -problem $(GRAPH_PROBLEM)

test:
	go test ./...
