package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func countingSort(as []int, bs []int, k int) {
	cs := make([]int, k+1)
	for _, x := range as {
		cs[x] = cs[x] + 1
	}
	for i := 1; i < k+1; i++ {
		cs[i] = cs[i] + cs[i-1]
	}
	for i := (len(as) - 1); i >= 0; i-- {
		bs[cs[as[i]]-1] = as[i]
		cs[as[i]] = cs[as[i]] - 1
	}
}

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}

func printArray(xs []int) {
	fmt.Println(strings.Trim(fmt.Sprint(xs), "[]"))
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	n := nextInt(sc)
	as := make([]int, n)
	k := 0
	for i := 0; i < n; i++ {
		x := nextInt(sc)
		as[i] = x
		if x > k {
			k = x
		}
	}
	bs := make([]int, n)
	countingSort(as, bs, k)
	printArray(bs)
}
