// Package main is for AOC day 5
package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

var (
	inputFile = flag.String("input-file", "input.data", "your input data file")
)

func init() {
	flag.Parse()
}

func fatal(format string, a ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", a...)
	os.Exit(1)
}

type stacks map[int][]byte

func fillStacks(s stacks, row []byte) {
	for i, c := range row {
		if (i+1)%2 == 0 && (i+1)%4 != 0 {
			if c >= 'A' && c <= 'Z' {
				s[i/4+1] = append([]byte{c}, s[i/4+1]...)
			}
		}
	}
}

func moveCrates(s stacks, row []byte, multiple bool) {
	c := bytes.Fields(row)
	move, _ := strconv.Atoi(string(c[1]))
	from, _ := strconv.Atoi(string(c[3]))
	to, _ := strconv.Atoi(string(c[5]))

	lf := len(s[from])

	if multiple {
		s[to] = append(s[to], s[from][lf-move:]...)
		s[from] = s[from][0 : lf-move]
	} else {
		for i := 0; i < move; i++ {
			s[to] = append(s[to], s[from][lf-1])
			s[from] = s[from][0 : lf-1]
			lf = len(s[from])
		}
	}
}

func getTopCrates(s stacks) (t []byte) {
	for i := 1; i <= len(s); i++ {
		t = append(t, s[i][len(s[i])-1])
	}

	return
}

func main() {
	f, err := os.ReadFile(*inputFile)
	if err != nil {
		fatal("could not read the file: %s", err)
	}

	start := time.Now()

	rows := bytes.Split(f, []byte("\n"))
	move := false

	s := stacks{}
	s9001 := stacks{}

	for _, row := range rows {
		l := len(row)

		if l != 0 && !move {
			fillStacks(s, row)
			fillStacks(s9001, row)
		} else if l == 0 && !move {
			move = true
		}

		if l != 0 && move {
			moveCrates(s, row, false)
			moveCrates(s9001, row, true)
		}
	}

	fmt.Printf("crates on top after rearrangement procedure: %s\n", getTopCrates(s))
	fmt.Printf("crates on top after rearrangement procedure for CrateMover 9001: %s\n", getTopCrates(s9001))
	fmt.Printf("\nstats:\nelapsed time: %d μs\n", time.Since(start).Microseconds())
}
