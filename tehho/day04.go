package main

import (
	"strconv"
	"strings"
)

func ParseStringArrayToInt(s []string) []int {
	val := make([]int, len(s))

	for i, v := range s {
		temp, err := strconv.Atoi(v)

		if err != nil {
			panic(err)
		}

		val[i] = temp
	}

	return val
}

func Day04part01(input string) int {
	rows := strings.Split(input, "\n")

	score := 0

	for _, row := range rows {
		temp := strings.Split(row, ",")
		elf1 := ParseStringArrayToInt(strings.Split(temp[0], "-"))
		elf2 := ParseStringArrayToInt(strings.Split(temp[1], "-"))

		if (elf1[0] <= elf2[0] && elf2[1] <= elf1[1]) || (elf2[0] <= elf1[0] && elf1[1] <= elf2[1]) {
			score++
		}

	}

	return score
}

func Day04part02(input string) int {
	rows := strings.Split(input, "\n")

	score := 0

	for _, row := range rows {
		temp := strings.Split(row, ",")
		elf1 := ParseStringArrayToInt(strings.Split(temp[0], "-"))
		elf2 := ParseStringArrayToInt(strings.Split(temp[1], "-"))

		if (elf1[0] <= elf2[0] && elf2[0] <= elf1[1]) || (elf1[0] <= elf2[1] && elf2[1] <= elf1[1]) || (elf2[0] <= elf1[0] && elf1[0] <= elf2[1]) {
			score++
		}

	}

	return score
}
