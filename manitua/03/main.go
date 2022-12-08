// Package main is for AOC day 3
package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"time"
)

var (
	inputFile = flag.String("input-file", "input.data", "your input data file")
)

func init() {
	flag.Parse()
}

func error(format string, a ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", a...)
	os.Exit(1)
}

func sum(a []int) int {
	var r int
	for _, v := range a {
		r += v
	}

	return r
}

func commonItems(c1, c2 []byte) (c []byte) {
	i := make(map[byte]bool)
	for _, v := range c1 {
		i[v] = true
	}

	for _, v := range c2 {
		if _, ok := i[v]; ok {
			c = append(c, v)
		}
	}

	return
}

func priority(i byte) int {
	if i >= 'a' && i <= 'z' {
		return int(i - 0x60) // 1 - 26
	} else if i >= 'A' && i <= 'Z' {
		return int(i - 0x26) // 27 - 52
	}

	return 0
}

func main() {
	f, err := os.ReadFile(*inputFile)
	if err != nil {
		error("could not read the file: %s", err)
	}

	start := time.Now()

	rs := bytes.Split(f, []byte("\n"))

	s := [][]int{{}, {}}

	for i, r := range rs {
		l := len(r)
		if l != 0 {
			c1 := r[0 : l/2]
			c2 := r[l/2:]
			c := commonItems(c1, c2)
			s[0] = append(s[0], priority(c[0]))

			if (i+1)%3 == 0 {
				rc := commonItems(commonItems(rs[i], rs[i-1]), rs[i-2])
				s[1] = append(s[1], priority(rc[0]))
			}
		}
	}

	fmt.Printf("sum of priorities: %d\n", sum(s[0]))
	fmt.Printf("sum of priorities per three-Elf group: %d\n", sum(s[1]))
	fmt.Printf("\nstats:\nelapsed time: %d μs\n", time.Since(start).Microseconds())
}
