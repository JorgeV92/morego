package main

import (
	"bufio"
	"fmt"
	"os"

	"morego/algorithms/data"
)

func main() {
	in := bufio.NewReaderSize(os.Stdin, 1<<20)
	out := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	edges := make([]data.Edge, 0, m)
	for i := 0; i < m; i++ {
		var from, to, cost int
		fmt.Fscan(in, &from, &to, &cost)

		edges = append(edges, data.Edge{
			U: from - 1,
			V: to - 1,
			W: cost,
		})
	}

	totalCost, mst := data.KruskalDSU(n, edges)
	if len(mst) != n-1 {
		fmt.Fprintln(out, "IMPOSSIBLE")
		return
	}

	fmt.Fprintln(out, totalCost)
}
