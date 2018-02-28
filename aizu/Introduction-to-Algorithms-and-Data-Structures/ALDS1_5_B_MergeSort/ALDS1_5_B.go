package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const max = 1000000001

var count = 0

func merge(as []int, bs []int) []int {
	var result []int
	as = append(as, max)
	bs = append(bs, max)
	i := 0
	j := 0
	for {
		if i == len(as)-1 && j == len(bs)-1 {
			break
		}
		count++
		if as[i] < bs[j] {
			if as[i] < max {
				result = append(result, as[i])
				i++
			}
		} else {
			if bs[j] < max {
				result = append(result, bs[j])
				j++
			}
		}
	}
	return result
}

func mergeSort(xs []int) []int {
	if len(xs) == 1 {
		return xs
	}
	m := len(xs) / 2
	var as = make([]int, len(xs[:m]))
	copy(as, xs[:m])
	var bs = make([]int, len(xs[m:]))
	copy(bs, xs[m:])
	return merge(mergeSort(as), mergeSort(bs))
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
	xs := make([]int, n)
	for i := 0; i < n; i++ {
		xs[i] = nextInt(sc)
	}
	result := mergeSort(xs)
	printArray(result)
	fmt.Println(count)
}
