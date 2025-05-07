package main

import (
	"bufio"
	"codeforces-tut/segtree/part3/b"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	s, _ := reader.ReadString('\n')
	parts := strings.Fields(s)
	n, _ := strconv.Atoi(parts[0])

	nums := make([]int, n)
	s, _ = reader.ReadString('\n')
	parts = strings.Fields(s)
	for i := 0; i < n; i++ {
		nums[i], _ = strconv.Atoi(parts[i])
	}

	arr := make([]any, n)

	st := b.NewSegTree(n)
	for i := n - 1; i >= 0; i-- {
		pos := st.Query(nums[i] + 1)
		st.Update(pos)
		arr[i] = pos + 1
	}

	fmt.Fprintln(writer, arr...)
}
