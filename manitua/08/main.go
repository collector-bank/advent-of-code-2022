// Package main is for AOC day 8
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

func fatal(format string, a ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", a...)
	os.Exit(1)
}

type layout [][]byte

func mapData(data [][]byte) (l layout) {
	for _, row := range data {
		if len(row) != 0 {
			r := []byte{}
			for _, v := range row {
				r = append(r, v)
			}
			l = append(l, r)
		}
	}

	return
}

func visibleTrees(lo layout) (vTree int) {
	ln := len(lo)
	for row, rValue := range lo {
		if row == 0 || row == ln-1 {
			vTree += ln
			continue
		}

		for col, cValue := range rValue {
			if col == 0 || col == ln-1 {
				vTree++
				continue
			}

			// Right
			visible := true
			for c := col + 1; c < ln; c++ {
				if rValue[c] >= cValue {
					visible = false
					break
				}
			}
			if visible {
				vTree++
				continue
			}

			// Left
			visible = true
			for c := 0; c < col; c++ {
				if rValue[c] >= cValue {
					visible = false
					break
				}
			}
			if visible {
				vTree++
				continue
			}

			// Down
			visible = true
			for r := row + 1; r < ln; r++ {
				if lo[r][col] >= cValue {
					visible = false
					break
				}
			}
			if visible {
				vTree++
				continue
			}

			// Up
			visible = true
			for r := 0; r < row; r++ {
				if lo[r][col] >= cValue {
					visible = false
					break
				}
			}
			if visible {
				vTree++
				continue
			}
		}
	}

	return
}

func maxScenicScore(lo layout) (mScore int) {
	ln := len(lo)

	for row, rValue := range lo {
		for col, cValue := range rValue {
			score := 0

			if row == 0 || row == ln-1 || col == 0 || col == ln-1 {
				continue
			}

			right := 0
			for c := col + 1; c < ln; c++ {
				if rValue[c] >= cValue || c == ln-1 {
					right = c - col
					break
				}
			}

			left := 0
			for c := col - 1; c >= 0; c-- {
				if rValue[c] >= cValue || c == 0 {
					left = col - c
					break
				}
			}

			down := 0
			for r := row + 1; r < ln; r++ {
				if lo[r][col] >= cValue || r == ln-1 {
					down = r - row
					break
				}
			}

			up := 0
			for r := row - 1; r >= 0; r-- {
				if lo[r][col] >= cValue || r == 0 {
					up = row - r
					break
				}
			}

			score = right * left * down * up

			if score > mScore {
				mScore = score
			}
		}
	}

	return
}

func main() {
	f, err := os.ReadFile(*inputFile)
	if err != nil {
		fatal("could not read the file: %s", err)
	}

	start := time.Now()

	l := mapData(bytes.Split(f, []byte("\n")))
	vt := visibleTrees(l)
	ms := maxScenicScore(l)

	fmt.Printf("number of trees visible from outside the grid: %d\n", vt)
	fmt.Printf("highest scenic score possible for any tree: %d\n", ms)
	fmt.Printf("\nstats:\nelapsed time: %d μs\n", time.Since(start).Microseconds())
}
