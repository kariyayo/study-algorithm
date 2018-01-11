package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Process struct {
	Name string
	Time int
}

type Queue struct {
	xs   [100001]*Process
	head int
	tail int
}

func (q *Queue) enqueue(x *Process) {
	q.xs[q.tail] = x
	if q.tail >= len(q.xs)-1 {
		q.tail = 0
	} else {
		q.tail++
	}
}

func (q *Queue) dequeue() *Process {
	x := q.xs[q.head]
	if q.head >= len(q.xs)-1 {
		q.head = 0
	} else {
		q.head++
	}
	return x
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
	q := nextInt(sc)

	var queue Queue
	for i := 0; i < n; i++ {
		queue.enqueue(&Process{
			Name: nextStr(sc),
			Time: nextInt(sc),
		})
	}

	sum := 0
	for n > 0 {
		p := queue.dequeue()
		if p.Time > q {
			sum = sum + q
			p.Time = p.Time - q
			queue.enqueue(p)
		} else {
			sum = sum + p.Time
			n--
			fmt.Print(p.Name)
			fmt.Print(" ")
			fmt.Println(sum)
		}
	}
}
