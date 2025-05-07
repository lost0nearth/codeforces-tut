package c

type SegTree struct {
	tree []int
	n    int
}

func (st *SegTree) update(node, lo, hi, i int) {
	if lo+1 == hi {
		st.tree[node] = 1
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

func (st *SegTree) query(node, lo, hi, l, r int) int {
	if r <= lo || hi <= l {
		return 0
	}
	if l <= lo && hi <= r {
		return st.tree[node]
	}
	mid := (lo + hi) / 2
	left := st.query(2*node+1, lo, mid, l, r)
	right := st.query(2*node+2, mid, hi, l, r)
	return left + right
}

func (st *SegTree) Update(idx int) {
	st.update(0, 0, st.n, idx)
}

func (st *SegTree) Query(l, r int) int {
	return st.query(0, 0, st.n, l, r)
}

func NewSegTree(n int) *SegTree {
	st := &SegTree{
		tree: make([]int, 4*n),
		n:    n,
	}
	return st
}
