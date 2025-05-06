package c

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
	st.tree[node] = max(st.tree[2*node+1], st.tree[2*node+2])
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
	st.tree[node] = max(st.tree[2*node+1], st.tree[2*node+2])
}

func (st *SegTree) query(node, lo, hi, x int) int {
	if st.tree[node] < x {
		return -1
	}
	if lo+1 == hi {
		return lo
	}

	mid := (lo + hi) / 2
	if st.tree[2*node+1] >= x {
		return st.query(2*node+1, lo, mid, x)
	}
	return st.query(2*node+2, mid, hi, x)
}

func (st *SegTree) Update(idx, val int) {
	st.update(0, 0, st.n, idx, val)
}

func (st *SegTree) Query(x int) int {
	return st.query(0, 0, st.n, x)
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
