package main

import (
	"strconv"
	"strings"
)

func FindMaxInArray(a []int) int {
	t := 0
	for _, i := range a {
		if i > t {
			t = i
		}
	}
	return t
}

func FindSumOf3Answers(input []int) int {
	a := []int{0, 0, 0}

	for _, i := range input {
		if i > a[0] {
			a[2] = a[1]
			a[1] = a[0]
			a[0] = i
		} else if i > a[1] {
			a[2] = a[1]
			a[1] = i
		} else if i > a[2] {
			a[2] = i
		}
	}

	return a[0] + a[1] + a[2]
}

func Day01part01(input string) int {
	rows := strings.Split(input, "\n")
	temp := []string{}
	answers := []int{}
	for _, row := range rows {
		if row == "" {
			a := 0

			for _, t := range temp {

				i, err := strconv.Atoi(t)

				if err != nil {
					panic(err)
				}

				a += i
			}

			answers = append(answers, a)
			temp = []string{}
		} else {
			temp = append(temp, row)
		}
	}

	a := 0

	for _, t := range temp {

		i, err := strconv.Atoi(t)

		if err != nil {
			panic(err)
		}

		a += i
	}

	answers = append(answers, a)

	return FindMaxInArray(answers)
}

func Day01part02(input string) int {
	rows := strings.Split(input, "\n")
	temp := []string{}
	answers := []int{}
	for _, row := range rows {
		if row == "" {
			a := 0

			for _, t := range temp {

				i, err := strconv.Atoi(t)

				if err != nil {
					panic(err)
				}

				a += i
			}

			answers = append(answers, a)
			temp = []string{}
		} else {
			temp = append(temp, row)
		}
	}

	a := 0

	for _, t := range temp {

		i, err := strconv.Atoi(t)

		if err != nil {
			panic(err)
		}

		a += i
	}

	answers = append(answers, a)

	return FindSumOf3Answers(answers)
}
