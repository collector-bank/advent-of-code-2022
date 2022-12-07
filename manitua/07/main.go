// Package main is for AOC day 7
package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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

type tree map[string]int

func dirTree(data [][]byte) tree {
	p := []string{}
	t := tree{}

	for _, r := range data {
		if len(r) != 0 {
			c := strings.Fields(string(r))
			if c[0] == "$" {
				if c[1] == "cd" && c[2] != ".." {
					p = append(p, c[2])
				} else if c[1] == "cd" && c[2] == ".." {
					p = p[:len(p)-1]
				}
			} else if c[0] != "$" && c[0] != "dir" {
				s, _ := strconv.Atoi(c[0])
				for i := range p {
					t[strings.Join(p[:i+1], "/")] += s
				}
			}
		}
	}

	return t
}

func main() {
	f, err := os.ReadFile(*inputFile)
	if err != nil {
		fatal("could not read the file: %s", err)
	}

	start := time.Now()

	data := bytes.Split(f, []byte("\n"))
	t := dirTree(data)

	sum := 0
	for _, v := range t {
		if v < 100_000 {
			sum += v
		}
	}

	unused := 70_000_000 - t["/"]
	needed := 30_000_000 - unused
	dirSizes := []int{}
	for _, v := range t {
		if v >= needed {
			dirSizes = append(dirSizes, v)
		}
	}
	sort.Ints(dirSizes)

	fmt.Printf("sum of directories with a total size of at most 100000: %d\n", sum)
	fmt.Printf("smallest directory, if deleted, that will free enough space: %d\n", dirSizes[0])
	fmt.Printf("\nstats:\nelapsed time: %d μs\n", time.Since(start).Microseconds())
}
