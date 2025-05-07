package main

import (
	"bufio"
	"codeforces-tut/segtree/part3/a"
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

	s, _ = reader.ReadString('\n')
	parts = strings.Fields(s)

	st := a.NewSegTree(n + 1)
	for i := range n {
		num, _ := strconv.Atoi(parts[i])
		st.Update(num)
		fmt.Fprint(writer, st.Query(num+1), " ")
	}
}
