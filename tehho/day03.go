package main

import (
	"strings"
)

func FindFirstCommonValue(a string, b string) rune {
	for _, c := range a {
		for _, d := range b {
			if c == d {
				return c
			}
		}
	}
	return '0'
}

func FindAllCommonValue(a string, b string) string {
	val := ""
	for _, c := range a {
		for _, d := range b {
			if c == d {
				val = val + string(c)
			}
		}
	}
	return val
}

func Day03part01(input string) int {
	rows := strings.Split(input, "\n")

	score := 0

	for _, row := range rows {
		halfLength := len(row) / 2

		common := FindFirstCommonValue(row[0:halfLength], row[halfLength:])

		if 'a' <= common && common <= 'z' {
			score += int(common-'a') + 1
		} else if 'A' <= common && common <= 'Z' {
			score += int(common-'A') + 27
		}
	}

	return score
}

func Day03part02(input string) int {
	rows := strings.Split(input, "\n")

	score := 0

	for i := 0; i < len(rows); i += 3 {
		a := rows[i]
		b := rows[i+1]
		c := rows[i+2]

		temp := FindAllCommonValue(a, b)
		common := FindFirstCommonValue(temp, c)

		if 'a' <= common && common <= 'z' {
			score += int(common-'a') + 1
		} else if 'A' <= common && common <= 'Z' {
			score += int(common-'A') + 27
		} else {
			panic("Not possible")
		}
	}

	return score
}
