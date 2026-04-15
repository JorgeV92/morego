package data

import "math/rand"

type Node struct {
	pri  int
	val  int
	c    [2]*Node
	sz   int
	sum  int64
	flip bool
}

func NewNode(v int) *Node {
	return &Node{
		pri: rand.Int(),
		val: v,
		sz:  1,
		sum: int64(v),
	}
}

func GetSize(x *Node) int {
	if x == nil {
		return 0
	}
	return x.sz
}

func GetSum(x *Node) int64 {
	if x == nil {
		return 0
	}
	return x.sum
}

func Prop(x *Node) *Node {
	if x == nil || !x.flip {
		return x
	}
	x.c[0], x.c[1] = x.c[1], x.c[0]
	x.flip = false
	if x.c[0] != nil {
		x.c[0].flip = !x.c[0].flip
	}
	if x.c[1] != nil {
		x.c[1].flip = !x.c[1].flip
	}
	return x
}

func Calc(x *Node) *Node {
	if x == nil {
		return nil
	}
	Prop(x.c[0])
	Prop(x.c[1])
	x.sz = 1 + GetSize(x.c[0]) + GetSize(x.c[1])
	x.sum = int64(x.val) + GetSum(x.c[0]) + GetSum(x.c[1])
	return x
}

func Inorder(x *Node, out *[]int) {
	if x == nil {
		return
	}
	Prop(x)
	Inorder(x.c[0], out)
	*out = append(*out, x.val)
	Inorder(x.c[1], out)
}

func Split(t *Node, v int) (*Node, *Node) {
	if t == nil {
		return nil, nil
	}
	Prop(t)

	if t.val >= v {
		a, b := Split(t.c[0], v)
		t.c[0] = b
		return a, Calc(t)
	}

	a, b := Split(t.c[1], v)
	t.c[1] = a
	return Calc(t), b
}

func SplitSz(t *Node, sz int) (*Node, *Node) {
	if t == nil {
		return nil, nil
	}
	Prop(t)

	leftSize := GetSize(t.c[0])

	if leftSize >= sz {
		a, b := SplitSz(t.c[0], sz)
		t.c[0] = b
		return a, Calc(t)
	}

	a, b := SplitSz(t.c[1], sz-leftSize-1)
	t.c[1] = a
	return Calc(t), b
}

func Merge(l, r *Node) *Node {
	if l == nil {
		return r
	}
	if r == nil {
		return l
	}

	Prop(l)
	Prop(r)

	if l.pri > r.pri {
		l.c[1] = Merge(l.c[1], r)
		return Calc(l)
	}
	r.c[0] = Merge(l, r.c[0])
	return Calc(r)
}

func Insert(root *Node, v int) *Node {
	a, b := Split(root, v)
	_, c := Split(b, v+1)
	return Merge(a, Merge(NewNode(v), c))
}

func Erase(root *Node, v int) *Node {
	a, b := Split(root, v)
	_, c := Split(b, v+1)
	return Merge(a, c)
}
