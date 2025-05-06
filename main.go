package main

import (
	"bufio"
	"codeforces-tut/segtree/part2/a"
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

	st := a.NewSegTree(nums)
	fmt.Fprintln(writer, st.Query())

	for range m {
		s, _ = reader.ReadString('\n')
		parts = strings.Fields(s)
		i, _ := strconv.Atoi(parts[0])
		v, _ := strconv.Atoi(parts[1])

		st.Update(i, v)
		fmt.Fprintln(writer, st.Query())
	}
}
