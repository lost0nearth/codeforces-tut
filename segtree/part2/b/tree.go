package b

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
	st.tree[node] = st.tree[2*node+1] + st.tree[2*node+2]
}

func (st *SegTree) update(node, lo, hi, i int) {
	if lo+1 == hi {
		st.tree[node] ^= 1
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
	left := 2*node+1
	mid := (lo+hi)/2
	if st.tree[left] >= k {
		return st.query(2*node+1, lo, mid, k)
	}
	return st.query(2*node+2, mid, hi, k-st.tree[left])
}

func (st *SegTree) Update(idx int) {
	st.update(0, 0, st.n, idx)
}

func (st *SegTree) Query(k int) int {
	return st.query(0, 0, st.n, k)
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
