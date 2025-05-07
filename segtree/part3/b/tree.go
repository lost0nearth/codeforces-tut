package b

type SegTree struct {
	tree []int
	n    int
}

func (st *SegTree) build(node, lo, hi int) {
	if lo+1 == hi {
		st.tree[node] = 1
		return
	}
	mid := (lo + hi) / 2
	st.build(2*node+1, lo, mid)
	st.build(2*node+2, mid, hi)
	st.tree[node] = st.tree[2*node+1] + st.tree[2*node+2]
}

func (st *SegTree) update(node, lo, hi, i int) {
	if lo+1 == hi {
		st.tree[node] = 0
		return
	}
	mid := (lo + hi) / 2
	if i < mid {
		st.update(2*node+1, lo, mid, i)
	} else {
		st.update(2*node+2, mid, hi, i)
	}
	st.tree[node] = st.tree[2*node+1] + st.tree[2*node+2]
}

func (st *SegTree) query(node, lo, hi, k int) int {
	if lo+1 == hi {
		return lo
	}
	mid := (lo + hi) / 2
	if st.tree[2*node+2] >= k {
		return st.query(2*node+2, mid, hi, k)
	}
	return st.query(2*node+1, lo, mid, k-st.tree[2*node+2])
}

func (st *SegTree) Update(idx int) {
	st.update(0, 0, st.n, idx)
}

func (st *SegTree) Query(k int) int {
	return st.query(0, 0, st.n, k)
}

func NewSegTree(n int) *SegTree {
	st := &SegTree{
		tree: make([]int, 4*n),
		n:    n,
	}
	st.build(0, 0, st.n)
	return st
}
