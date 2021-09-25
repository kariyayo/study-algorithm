package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func distance(s, d Pair) float64 {
	result := math.Sqrt(math.Pow(float64(s.First-d.First), 2) + math.Pow(float64(s.Second-d.Second), 2))
	return result
}

func iter(d float64, p Pair, ps []Pair) []float64 {
	if len(ps) == 0 {
		return []float64{d}
	}
	result := []float64{}
	for i := 0; i < len(ps); i++ {
		child := ps[i]
		x := distance(p, child) + d
		left := ps[0:i]
		right := []Pair{}
		if i+1 <= len(ps) {
			right = ps[i+1:]
		}
		others := []Pair{}
		others = append(others, left...)
		others = append(others, right...)
		sums := iter(x, child, others)
		result = append(result, sums...)
	}
	return result
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0)
	sc.Buffer(buf, 1010000)
	sc.Split(bufio.ScanWords)

	N := nextInt(sc)
	towns := make([]Pair, N)
	for i := 0; i < N; i++ {
		x := nextInt(sc)
		y := nextInt(sc)
		towns[i] = Pair{x, y}
	}
	cnt := 0.0
	sum := 0.0
	for i := 0; i < len(towns); i++ {
		p := towns[i]
		left := towns[0:i]
		right := []Pair{}
		if i+1 <= len(towns) {
			right = towns[i+1:]
		}
		others := []Pair{}
		others = append(others, left...)
		others = append(others, right...)
		ds := iter(0.0, p, others)
		for _, d := range ds {
			cnt++
			sum += d
		}
	}
	fmt.Println(sum / cnt)
}

// ----------
type Pair struct {
	First  int
	Second int
}

func toInt(s string) int {
	x, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return x
}

func toInt64(s string) uint64 {
	x, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minUint64(a, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func nextString(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}

func nextNumber(sc *bufio.Scanner) float64 {
	sc.Scan()
	f, err := strconv.ParseFloat(sc.Text(), 32)
	if err != nil {
		panic(err)
	}
	return f
}

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}

func nextUint64(sc *bufio.Scanner) uint64 {
	sc.Scan()
	n, err := strconv.ParseUint(sc.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	return n
}

func printArray(xs []int) {
	fmt.Println(strings.Trim(fmt.Sprint(xs), "[]"))
}

func debugPrintArray(xs []int) {
	fmt.Fprintf(os.Stderr, "\n---\n%v\n", strings.Trim(fmt.Sprint(xs), "[]"))
}

func debugPrintf(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
}
