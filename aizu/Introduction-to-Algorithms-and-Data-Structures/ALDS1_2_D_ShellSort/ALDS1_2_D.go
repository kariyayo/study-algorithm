package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func insertionSort(n int, as []int, g int, count *int) {
	for i := 0; i < n-g; i++ {
		for j := i + g; j > g-1; j = j - g {
			if as[j] < as[j-g] {
				a := as[j]
				as[j] = as[j-g]
				as[j-g] = a
				*count = *count + 1
			} else {
				j = g - 1
			}
		}
	}
}

func shellSort(n int, as []int) {
	count := 0
	var gs []int
	for i := 0.0; i < 100 && n/int(math.Pow(2, i)) >= 1; i++ {
		gs = append(gs, n/int(math.Pow(2, i)))
	}
	m := len(gs)

	for _, g := range gs {
		insertionSort(n, as, g, &count)
	}

	fmt.Println(m)
	printArray(gs)
	fmt.Println(count)
	for _, v := range as {
		fmt.Println(v)
	}
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
	var as []int
	for i := 0; i < n; i++ {
		a := nextInt(sc)
		as = append(as, a)
	}
	shellSort(n, as)
}
