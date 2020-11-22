package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func maxHeapify(n int, as []int, i int) {
	l := 2*i + 1
	r := 2*i + 2
	largest := i
	if l < n && as[l] > as[largest] {
		largest = l
	}
	if r < n && as[r] > as[largest] {
		largest = r
	}
	if i != largest {
		tmp := as[i]
		as[i] = as[largest]
		as[largest] = tmp
		maxHeapify(n, as, largest)
	}
}

func buildMaxHeap(n int, as []int) {
	for i := (n - 1) / 2; i >= 0; i-- {
		maxHeapify(n, as, i)
	}
}

func debugPrintf(format string, a ...interface{}) {
	// fmt.Printf(format, a...)
}

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	n := nextInt(sc)
	as := make([]int, n)
	as[0] = 0
	for i := 0; i < n; i++ {
		as[i] = nextInt(sc)
	}
	buildMaxHeap(n, as)
	for i := 0; i < n; i++ {
		fmt.Printf(" %d", as[i])
	}
	fmt.Println()
}
