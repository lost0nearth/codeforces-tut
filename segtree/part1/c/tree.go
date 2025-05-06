package c

import "math"

type Node struct {
	val, count int
}

func (n Node) Val() int {
	return n.val
}

func (n Node) Count() int {
	return n.count
}

type SegTree struct {
	tree []Node
	n    int
}

func merge(n1, n2 Node) Node {
	if n1.val < n2.val {
		return n1
	} else if n1.val > n2.val {
		return n2
	}
	return Node{n1.val, n1.count + n2.count}
}

func (st *SegTree) build(node, lo, hi int, nums []int) {
	if lo+1 == hi {
		st.tree[node] = Node{nums[lo], 1}
		return
	}
	mid := (lo + hi) / 2
	st.build(2*node+1, lo, mid, nums)
	st.build(2*node+2, mid, hi, nums)
	st.tree[node] = merge(st.tree[2*node+1], st.tree[2*node+2])
}

func (st *SegTree) update(node, lo, hi, i, v int) {
	if lo+1 == hi {
		st.tree[node].val = v
		return
	}
	mid := (lo + hi) / 2
	if i < mid {
		st.update(2*node+1, lo, mid, i, v)
	} else {
		st.update(2*node+2, mid, hi, i, v)
	}
	st.tree[node] = merge(st.tree[2*node+1], st.tree[2*node+2])
}

func (st *SegTree) query(node, lo, hi, l, r int) Node {
	if r <= lo || hi <= l {
		return Node{math.MaxInt, 0}
	}
	if l <= lo && hi <= r {
		return st.tree[node]
	}
	mid := (lo + hi) / 2
	left := st.query(2*node+1, lo, mid, l, r)
	right := st.query(2*node+2, mid, hi, l, r)
	return merge(left, right)
}

func (st *SegTree) Update(idx, val int) {
	st.update(0, 0, st.n, idx, val)
}

func (st *SegTree) Query(l, r int) Node {
	return st.query(0, 0, st.n, l, r)
}

func NewSegTree(nums []int) *SegTree {
	n := len(nums)
	st := &SegTree{
		tree: make([]Node, 4*n),
		n:    n,
	}
	st.build(0, 0, st.n, nums)
	return st
}
