package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type hash struct {
	ss [1000000]string
}

func (h *hash) insert(data string) {
	k := toKey(data)
	i := 0
	hValue := h.hash(k, i)
	v := h.ss[hValue]
	for v != "" && v != data {
		i++
		hValue = h.hash(k, i)
		v = h.ss[hValue]
	}
	if h.ss[hValue] == "" {
		h.ss[hValue] = data
	}
}

func (h *hash) search(data string) bool {
	k := toKey(data)
	i := 0
	v := h.ss[h.hash(k, i)]
	for v != "" && v != data {
		i++
		v = h.ss[h.hash(k, i)]
	}
	if v == "" {
		return false
	}
	return true
}

func (h *hash) hash(k int, i int) int {
	return (h1(k) + i*h2(k)) % 1046527
}

func h1(k int) int {
	return k % 1046527
}

func h2(k int) int {
	return 1 + (k % 1046526)
}

func toKey(s string) int {
	result := 0
	for _, c := range s {
		switch c {
		case 'A':
			result = result + 1
		case 'C':
			result = result + 2
		case 'G':
			result = result + 3
		case 'T':
			result = result + 4
		default:
		}
	}
	return result
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

	var h hash

	n := nextInt(sc)
	for n > 0 {
		cmd := nextStr(sc)
		v := nextStr(sc)
		switch cmd {
		case "insert":
			h.insert(v)
		case "find":
			if h.search(v) {
				fmt.Println("yes")
			} else {
				fmt.Println("no")
			}
		default:
		}
		n--
	}
}
