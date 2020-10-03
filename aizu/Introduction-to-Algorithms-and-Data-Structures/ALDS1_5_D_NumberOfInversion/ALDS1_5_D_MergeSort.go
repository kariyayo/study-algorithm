package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ans int = 0

func merge(l, r []int) []int {
	n := len(l) + len(r)
	result := make([]int, n)

	li := 0
	ri := 0

	for k := 0; k < n; k++ {
		lv := 1000000000
		if li < len(l) {
			lv = l[li]
		}
		rv := 1000000000
		if ri < len(r) {
			rv = r[ri]
		}
		if lv > rv {
			// 左の方が大きいので置き換えが発生
			if li < len(l) {
				ans = ans + len(l) - li
			}

			result[k] = rv
			ri++
		} else {
			result[k] = lv
			li++
		}
	}

	return result
}

func mergeSort(xs []int) []int {
	if len(xs) == 0 {
		return []int{}
	} else if len(xs) == 1 {
		return []int{xs[0]}
	} else {
		m := len(xs) / 2
		l := mergeSort(xs[:m])
		r := mergeSort(xs[m:])
		return merge(l, r)
	}
}

func solve(n int, xs []int) {
	_ = mergeSort(xs)
	// fmt.Println(printArray(sorted))
}

func printArray(xs []int) string {
	return strings.Trim(fmt.Sprint(xs), "[]")
}

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	s := sc.Text()
	x, err := strconv.Atoi(s)
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
	solve(n, xs)
	fmt.Println(ans)
}
