package a

type Node struct {
	val      int
	sum      int
	pre, suf int
}

type SegTree struct {
	tree []Node
	n    int
}

func merge(n1, n2 Node) Node {
	var node Node
	node.pre = max(n1.pre, n1.sum+n2.pre)
	node.suf = max(n2.suf, n1.suf+n2.sum)
	node.val = max(n1.val, n2.val, n1.suf+n2.pre)
	node.sum = n1.sum + n2.sum
	return node
}

func (st *SegTree) build(node, lo, hi int, nums []int) {
	if lo+1 == hi {
		st.tree[node] = Node{
			pre: max(0, nums[lo]),
			suf: max(0, nums[lo]),
			val: max(0, nums[lo]),
			sum: nums[lo],
		}
		return
	}
	mid := (lo + hi) / 2
	st.build(2*node+1, lo, mid, nums)
	st.build(2*node+2, mid, hi, nums)
	st.tree[node] = merge(st.tree[2*node+1], st.tree[2*node+2])
}

func (st *SegTree) update(node, lo, hi, i, v int) {
	if lo+1 == hi {
		st.tree[node] = Node{
			pre: max(0, v),
			suf: max(0, v),
			val: max(0, v),
			sum: v,
		}
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

func (st *SegTree) Update(idx, val int) {
	st.update(0, 0, st.n, idx, val)
}

func (st *SegTree) Query() int {
	return st.tree[0].val
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
