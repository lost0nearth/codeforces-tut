package main

import (
	"bufio"
	"codeforces-tut/segtree/part3/c"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	arr := make([]int, 2*n)
	for i := range arr {
		fmt.Fscan(reader, &arr[i])
	}

	pos := make([]int, n+1)
	for i := range pos {
		pos[i] = -1
	}
	res := make([]any, n+1)

	st := c.NewSegTree(2*n)

	for i, v := range arr {
		if pos[v] == -1 {
			pos[v] = i
		} else {
			res[v] = st.Query(pos[v], i)
			st.Update(pos[v])
		}
	}

	fmt.Fprintln(writer, res[1:]...)
}
