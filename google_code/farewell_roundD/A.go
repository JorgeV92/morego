package main

import (
	"bufio"
	"fmt"
	"os"
)

type FastScanner struct {
	r *bufio.Reader
}

func NewFastScanner() *FastScanner {
	return &FastScanner{r: bufio.NewReaderSize(os.Stdin, 1<<20)}
}

func (fs *FastScanner) NextInt() int {
	sign, val := 1, 0
	c, err := fs.r.ReadByte()
	for (c < '0' || c > '9') && c != '-' {
		c, err = fs.r.ReadByte()
		if err != nil {
			return 0
		}
	}
	if c == '-' {
		sign = -1
		c, _ = fs.r.ReadByte()
	}
	for c >= '0' && c <= '9' {
		val = val*10 + int(c-'0')
		c, err = fs.r.ReadByte()
		if err != nil {
			break
		}
	}
	if err == nil {
		_ = fs.r.UnreadByte()
	}
	return sign * val
}

type SideData struct {
	n   int
	loc []int64 // loc[v] = sum of distances from v to all nodes on this size
	tot int64   // sum of distances over all unordered pairs on this side
}

func buildSide(n int, fs *FastScanner) SideData {
	adj := make([][]int, n)
	for v := 0; v+1 < n; v++ {
		u := fs.NextInt() - 1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}
	parent := make([]int, n)
	for i := range parent {
		parent[i] = -2
	}
	parent[0] = -1
	order := make([]int, 0, n)
	stack := []int{0}

	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		order = append(order, cur)

		for _, nxt := range adj[cur] {
			if nxt == parent[cur] {
				continue
			}
			parent[nxt] = cur
			stack = append(stack, nxt)
		}
	}

	sz := make([]int, n)
	loc := make([]int64, n)

	var rootDistSum int64
	for i := n - 1; i >= 0; i-- {
		v := order[i]
		sz[v] = 1
		for _, nxt := range adj[v] {
			if parent[nxt] == v {
				sz[v] += sz[nxt]
				rootDistSum += int64(sz[nxt])
			}
		}
	}
	loc[0] = rootDistSum

	// reroot DP
	for _, v := range order {
		for _, nxt := range adj[v] {
			if parent[nxt] == v {
				loc[nxt] = loc[v] - int64(sz[nxt]) + int64(n-sz[nxt])
			}
		}
	}

	// sum of all pairwise distances inside this side
	var total int64
	for _, x := range loc {
		total += x
	}
	total /= 2

	return SideData{
		n:   n,
		loc: loc,
		tot: total,
	}
}

func main() {
	in := NewFastScanner()
	out := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer out.Flush()

	T := in.NextInt()
	// time O(W + E + C)
	// space O(W * E)
	for tc := 1; tc <= T; tc++ {
		W := in.NextInt()
		E := in.NextInt()
		C := in.NextInt()

		west := buildSide(W, in)
		east := buildSide(E, in)

		totalNodes := int64(W + E)
		denom := totalNodes * (totalNodes - 1) / 2

		fmt.Fprintf(out, "Case #%d:", tc)
		for q := 0; q < C; q++ {
			a := in.NextInt() - 1 // west
			b := in.NextInt() - 1 // east
			var total int64
			total += west.tot
			total += east.tot
			total += west.loc[a] * int64(E)
			total += int64(W * E)
			total += east.loc[b] * int64(W)

			avg := float64(total) / float64(denom)
			fmt.Fprintf(out, " %.12f", avg)
		}
		fmt.Fprintln(out)
	}
}
