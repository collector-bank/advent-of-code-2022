// Package main is for AOC day 4
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

func arrangeRanges(ranges *[][]byte) ([]int, error) {
	c := []int{}
	for _, i := range *ranges {
		rv := bytes.Split(i, []byte("-"))
		for _, j := range rv {
			n, err := strconv.Atoi(string(j))
			if err != nil {
				return c, fmt.Errorf("invalid range: %s", err)
			}
			c = append(c, n)
		}
	}

	return c, nil
}

func compareRanges(ranges *[][]byte) (bool, bool, error) {
	c, err := arrangeRanges(ranges)
	if err != nil {
		return false, false, fmt.Errorf("unable to arrange ranges: %s", err)
	}

	if (c[0] <= c[2] && c[1] >= c[3]) || (c[2] <= c[0] && c[3] >= c[1]) { // A ⊆ B or B ⊆ A
		return true, false, nil
	} else if c[0] <= c[3] && c[2] <= c[1] { // A ∩ B
		return false, true, nil
	}

	return false, false, nil
}

func main() {
	f, err := os.ReadFile(*inputFile)
	if err != nil {
		fatal("could not read the file: %s", err)
	}

	start := time.Now()

	ps := bytes.Split(f, []byte("\n"))

	s := 0
	i := 0

	for _, p := range ps {
		l := len(p)
		if l != 0 {
			ap := bytes.Split(p, []byte(","))
			c, o, err := compareRanges(&ap)
			if err != nil {
				fatal("unable to compare ranges: %s", err)
			}

			if c {
				s++
			}

			if o {
				i++
			}
		}
	}

	fmt.Printf("number of assignment pairs where range fully contain the other: %d\n", s)
	fmt.Printf("number of assignment pairs where ranges overlap: %d\n", s+i)
	fmt.Printf("\nstats:\nelapsed time: %d μs\n", time.Since(start).Microseconds())
}
