// Package main is for AOC day 6
package main

import (
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

func fatal(format string, a ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", a...)
	os.Exit(1)
}

func findMarker(data []byte, mLength int) int {
	for i := range data {
		if i >= mLength-1 {
			d := data[i-mLength+1 : i+1]
			f := 0
			for j := range d {
				for k := range d {
					if d[j] == d[k] {
						f++
					}
				}
			}

			if f == mLength {
				return i + 1
			}
		}
	}

	return 0
}

func main() {
	f, err := os.ReadFile(*inputFile)
	if err != nil {
		fatal("could not read the file: %s", err)
	}

	start := time.Now()

	var m4, m14 int
	m4 = findMarker(f, 4)
	m14 = findMarker(f, 14)

	fmt.Printf("first start-of-packet marker (4 distinct characters): %d\n", m4)
	fmt.Printf("first start-of-packet marker (14 distinct characters): %d\n", m14)
	fmt.Printf("\nstats:\nelapsed time: %d μs\n", time.Since(start).Microseconds())
}
