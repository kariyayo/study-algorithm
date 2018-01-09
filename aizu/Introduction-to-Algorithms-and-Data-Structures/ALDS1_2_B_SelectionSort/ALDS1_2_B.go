package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func SelectionSort(n int, as []int) {
	count := 0
	for i := 0; i < n; i++ {
		min := as[i]
		min_index := i
		for j := i + 1; j < n; j++ {
			if min > as[j] {
				min = as[j]
				min_index = j
			}
		}
		if min_index > i {
			a := as[min_index]
			as[min_index] = as[i]
			as[i] = a
			count = count + 1
		}
	}
	printArray(as)
	fmt.Println(count)
}

func printArray(as []int) {
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

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	n := nextInt(sc)
	as := []int{}
	for i := 0; i < n; i++ {
		x := nextInt(sc)
		as = append(as, x)
	}
	SelectionSort(n, as)
}
