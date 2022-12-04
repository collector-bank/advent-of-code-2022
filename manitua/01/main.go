// Package main is for AOC day 1
package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
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

func main() {
	f, err := os.ReadFile(*inputFile)
	if err != nil {
		error("could not read the file: %s", err)
	}

	start := time.Now()

	s := bytes.Split(f, []byte("\n"))

	var cal int
	elfs := []int{}

	for _, l := range s {
		if len(l) != 0 {
			c, err := strconv.Atoi(string(l))
			if err != nil {
				error("invalid input: %s", err)
			}

			cal += c
		} else {
			elfs = append(elfs, cal)
			cal = 0
		}
	}

	sort.Ints(elfs)

	fmt.Printf("elf that carries most calories: %d\n", elfs[len(elfs)-1])
	fmt.Printf("calories carried by top three elf's: %d\n", sum(elfs[len(elfs)-3:]))
	fmt.Printf("\nstats:\nelapsed time: %d μs\n", time.Since(start).Microseconds())
}
