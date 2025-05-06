package b

import "math"

type SegTree struct {
	tree []int
	n    int
}

func (st *SegTree) build(node, lo, hi int, nums []int) {
	if lo+1 == hi {
		st.tree[node] = nums[lo]
		return
	}
	mid := (lo + hi) / 2
	st.build(2*node+1, lo, mid, nums)
	st.build(2*node+2, mid, hi, nums)
	st.tree[node] = min(st.tree[2*node+1], st.tree[2*node+2])
}

func (st *SegTree) update(node, lo, hi, i, v int) {
	if lo+1 == hi {
		st.tree[node] = v
		return
	}
	mid := (lo + hi) / 2
	if i < mid {
		st.update(2*node+1, lo, mid, i, v)
	} else {
		st.update(2*node+2, mid, hi, i, v)
	}
	st.tree[node] = min(st.tree[2*node+1], st.tree[2*node+2])
}

func (st *SegTree) query(node, lo, hi, l, r int) int {
	if r <= lo || hi <= l {
		return math.MaxInt
	}
	if l <= lo && hi <= r {
		return st.tree[node]
	}
	mid := (lo + hi) / 2
	left := st.query(2*node+1, lo, mid, l, r)
	right := st.query(2*node+2, mid, hi, l, r)
	return min(left, right)
}

func (st *SegTree) Update(idx, val int) {
	st.update(0, 0, st.n, idx, val)
}

func (st *SegTree) Query(l, r int) int {
	return st.query(0, 0, st.n, l, r)
}

func NewSegTree(nums []int) *SegTree {
	n := len(nums)
	st := &SegTree{
		tree: make([]int, 4*n),
		n:    n,
	}
	st.build(0, 0, st.n, nums)
	return st
}
