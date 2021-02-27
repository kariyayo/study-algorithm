package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	K := nextInt(sc)
	S := nextString(sc)[:4]
	T := nextString(sc)[:4]

	cards := make([]int, 10)
	for i := 0; i < 10; i++ {
		cards[i] = K
	}
	cards[0] = 0

	for i := 0; i < len(S); i++ {
		x, _ := strconv.Atoi(string(S[i]))
		cards[x]--
	}
	for i := 0; i < len(T); i++ {
		x, _ := strconv.Atoi(string(T[i]))
		cards[x]--
	}

	ok := 0
	for i := 1; i <= 9; i++ {
		scoreS := score(S, i)
		for j := 1; j <= 9; j++ {
			scoreT := score(T, j)
			if scoreS <= scoreT {
				continue
			}
			if i == j {
				ok += cards[i] * (cards[i] - 1)
			} else {
				ok += cards[i] * (cards[j])
			}
		}
	}

	// 全パターン = 残りのカード（9*K-(4*2)）を2人にそれぞれ1枚ずつ配るパターン（順列）
	total := (9*K - 8) * (9*K - 9)

	fmt.Println(float64(ok) / float64(total))
}

func score(S string, x int) int {
	s := S + fmt.Sprintf("%d", x)
	score := 0
	for i := 1; i <= 9; i++ {
		ci := strings.Count(s, fmt.Sprintf("%d", i))
		score += i * pow(10, ci)
	}
	return score
}

// ----------

func min(a, b int) int {
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

func debugPrintf(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
}
