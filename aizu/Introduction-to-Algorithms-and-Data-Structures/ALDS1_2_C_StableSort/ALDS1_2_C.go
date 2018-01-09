package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func bubbleSort(n int, as []string) {
	for i := range as {
		for j := n - 1; j > i; j-- {
			if as[j][1:] < as[j-1][1:] {
				a := as[j]
				as[j] = as[j-1]
				as[j-1] = a
			}
		}
	}
	printArray(as)
}

func selectionSort(n int, as []string) {
	for i := range as {
		min := as[i]
		min_index := i
		for j := i + 1; j < n; j++ {
			if min[1:] > as[j][1:] {
				min = as[j]
				min_index = j
			}
		}
		if min_index > i {
			a := as[min_index]
			as[min_index] = as[i]
			as[i] = a
		}
	}
	printArray(as)
}

func stableSort(n int, as []string) {
	bs := make([]string, n)
	copy(bs, as)
	bubbleSort(n, bs)
	fmt.Println("Stable")
	selectionSort(n, as)
	if isStable(as, bs) {
		fmt.Println("Stable")
	} else {
		fmt.Println("Not stable")
	}
}

func isStable(as []string, bs []string) bool {
	for i, a := range as {
		if a != bs[i] {
			return false
		}
	}
	return true
}

func printArray(as []string) {
	fmt.Println(strings.Trim(fmt.Sprint(as), "[]"))
}

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}

func nextStr(sc *bufio.Scanner) string {
	sc.Scan()
	s := sc.Text()
	return s
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	n := nextInt(sc)
	var as []string
	for i := 0; i < n; i++ {
		a := nextStr(sc)
		as = append(as, a)
	}
	stableSort(n, as)
}
