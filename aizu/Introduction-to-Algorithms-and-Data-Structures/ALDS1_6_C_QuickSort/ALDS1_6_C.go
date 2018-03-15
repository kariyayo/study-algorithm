package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type card struct {
	label string
	value int
}

// ----- Merge Sort -----
const max = 1000000001

var count = 0

func merge(as []card, bs []card) []card {
	var result []card
	as = append(as, card{value: max})
	bs = append(bs, card{value: max})
	i := 0
	j := 0
	for {
		if i == len(as)-1 && j == len(bs)-1 {
			break
		}
		count++
		if as[i].value <= bs[j].value {
			if as[i].value < max {
				result = append(result, as[i])
				i++
			}
		} else {
			if bs[j].value < max {
				result = append(result, bs[j])
				j++
			}
		}
	}
	return result
}

func mergeSort(xs []card) []card {
	if len(xs) == 1 {
		return xs
	}
	m := len(xs) / 2
	var as = make([]card, len(xs[:m]))
	copy(as, xs[:m])
	var bs = make([]card, len(xs[m:]))
	copy(bs, xs[m:])
	return merge(mergeSort(as), mergeSort(bs))
}

// ----------

// ----- Quick Sort -----
func partition(xs []card, p int, r int) int {
	x := xs[r]
	as := xs[:r]
	i := p - 1 // 前パーティションの最後の要素のポインタ
	j := p
	for j < len(as) {
		target := as[j]
		if target.value <= x.value {
			i++
			as[j] = as[i]
			as[i] = target
		}
		j++
	}
	xs[r] = xs[i+1]
	xs[i+1] = x
	return i + 1
}

func quickSort(l []card, p int, r int) {
	if p < r {
		q := partition(l, p, r)
		quickSort(l, p, q-1)
		quickSort(l, q+1, r)
	}
}

// ----------

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
	return sc.Text()
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	n := nextInt(sc)
	l := make([]card, n)
	for i := 0; i < n; i++ {
		l[i] = card{
			label: nextStr(sc),
			value: nextInt(sc),
		}
	}
	origin := make([]card, n)
	copy(origin, l)

	quickSort(l, 0, n-1)
	l2 := mergeSort(origin)

	count := 0
	for i := range l {
		if l[i] == l2[i] {
			count++
		} else {
			break
		}
	}
	if count == n {
		fmt.Println("Stable")
	} else {
		fmt.Println("Not stable")
	}
	for _, x := range l {
		fmt.Printf("%s %d\n", x.label, x.value)
	}
}
