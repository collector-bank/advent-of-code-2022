// Package main is for AOC day 2
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

func shapeScore(s byte) int {
	v := map[byte]int{
		'A': 1, // Rock
		'B': 2, // Paper
		'C': 3, // Scissors
	}
	return v[s]
}

func outcomeScore(elf, me byte) (s int) {
	switch elf - me {
	// (255) Elf: rock     - Me: paper    (A - B)
	// (255) Elf: paper    - Me: scissors (B - C)
	// (2)   Elf: scissors - Me: rock (C - A)
	case 255, 2:
		s = 6

	// (254) Elf: rock     - Me: scissors (A - C)
	// (1)   Elf: paper    - Me: rock (B - A)
	// (1)   Elf: scissors - Me: paper (C - B)
	case 254, 1:
		s = 0

	// Draw
	case 0:
		s = 3
	}

	return
}

func howToEndRound(elf, outcome byte) (c byte) {
	o := map[string]map[byte]byte{
		"lose": {
			'A': 'C', // Rock(A) elf wins
			'B': 'A', // Paper(B) elf wins
			'C': 'B', // Scissors(C) elf wins
		},
		"win": {
			'A': 'B', // Rock(A) elf loses
			'B': 'C', // Paper(B) elf loses
			'C': 'A', // Scissors(C) elf loses
		},
	}
	switch outcome {
	case 'X':
		c = o["lose"][elf]
	case 'Y':
		c = elf
	case 'Z':
		c = o["win"][elf]
	}

	return
}

func main() {
	f, err := os.ReadFile(*inputFile)
	if err != nil {
		error("could not read the file: %s", err)
	}

	start := time.Now()

	guide := bytes.Split(f, []byte("\n"))

	scores := [][]int{{}, {}}

	for _, line := range guide {
		if len(line) != 0 {
			round := bytes.Split(line, []byte(" "))

			elf := round[0][0]
			myReason := round[1][0] - 0x17 // Move X,Y,Z to correspond to A,B,C

			outcome := round[1][0]
			my := howToEndRound(elf, outcome)

			scores[0] = append(scores[0], outcomeScore(elf, myReason)+shapeScore(myReason))
			scores[1] = append(scores[1], outcomeScore(elf, my)+shapeScore(my))
		}
	}

	fmt.Printf("my total score (my reason): %d\n", sum(scores[0]))
	fmt.Printf("my total score (decrypted): %d\n", sum(scores[1]))
	fmt.Printf("\nstats:\nelapsed time: %d μs\n", time.Since(start).Microseconds())
}
