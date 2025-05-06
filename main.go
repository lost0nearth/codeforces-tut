package main

import (
	"bufio"
	"codeforces-tut/segtree/part2/c"
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
		first, _ := strconv.Atoi(parts[0])

		if first == 1 {
			idx, _ := strconv.Atoi(parts[1])
			val, _ := strconv.Atoi(parts[2])
			st.Update(idx, val)
		} else {
			x, _ := strconv.Atoi(parts[1])
			fmt.Fprintln(writer, st.Query(x))
		}
	}
}
