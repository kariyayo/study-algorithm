package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solver(n int, xs []int) {
	for i := 0; i < n; i++ {
		id := i + 1
		key := xs[i]
		ss := []string{fmt.Sprintf("node %d: key = %d", id, key)}
		if i != 0 && (i-1)/2 >= 0 {
			ss = append(ss, fmt.Sprintf(" parent key = %d", xs[(i-1)/2]))
		}
		if (2*i + 1) <= n-1 {
			ss = append(ss, fmt.Sprintf(" left key = %d", xs[2*i+1]))
		}
		if (2*i + 2) <= n-1 {
			ss = append(ss, fmt.Sprintf(" right key = %d", xs[2*i+2]))
		}
		ss = append(ss, " \n")
		fmt.Printf(strings.Join(ss, ","))
	}
}

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	x, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return x
}
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	n := nextInt(sc)
	xs := make([]int, n)
	for i := 0; i < n; i++ {
		xs[i] = nextInt(sc)
	}
	solver(n, xs)
}
