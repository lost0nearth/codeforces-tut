package main

import (
	"bufio"
	"codeforces-tut/segtree/part1/c"
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
	m, _ := strconv.Atoi(parts[1])

	nums := make([]int, n)
	s, _ = reader.ReadString('\n')
	parts = strings.Fields(s)

	for i := range nums {
		nums[i], _ = strconv.Atoi(parts[i])
	}

	st := c.NewSegTree(nums)

	for range m {
		s, _ = reader.ReadString('\n')
		parts = strings.Fields(s)
		if parts[0] == "1" {
			i, _ := strconv.Atoi(parts[1])
			v, _ := strconv.Atoi(parts[2])

			st.Update(i, v)
		} else {
			l, _ := strconv.Atoi(parts[1])
			r, _ := strconv.Atoi(parts[2])

			res := st.Query(l, r)
			fmt.Fprintln(writer, res.Val(), res.Count())
		}
	}
}
