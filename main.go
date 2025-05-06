package main

import (
	"bufio"
	"codeforces-tut/segtree/part2/b"
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

	st := b.NewSegTree(nums)


	for range m {
		s, _ = reader.ReadString('\n')
		parts = strings.Fields(s)
		first, _ := strconv.Atoi(parts[0])
		second, _ := strconv.Atoi(parts[1])

		if first == 1 {
			st.Update(second)
		} else {
			fmt.Fprintln(writer, st.Query(second+1))
		}
	}
}
