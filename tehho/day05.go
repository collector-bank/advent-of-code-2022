package main

import (
	"errors"
	"strconv"
	"strings"
)

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func min(i1, i2 int) int {
	if i1 < i2 {
		return i1
	}
	return i2
}

func (s *Stack) Pop() (string, error) {
	if s.IsEmpty() {
		return "", errors.New("Stack is empty")
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, nil
	}
}

func SplitN(str string, n int) []string {
	var val Stack
	for i := 0; i < len(str); i += n {
		val.Push(str[i:min(i+n, len(str))])
	}

	return val
}

func Day05part01(input string) string {
	temp := strings.Split(input, "\n\n")

	rows := strings.Split(temp[0], "\n")

	stacks := make([]Stack, (len(rows[0])/4)+1)

	for i := len(rows) - 1; i >= 0; i-- {
		content := SplitN(rows[i], 4)
		for j, c := range content {
			if c[2] != ' ' {
				stacks[j].Push(string(c[1]))
			}
		}
	}

	rows = strings.Split(temp[1], "\n")

	for _, row := range rows {
		// move 1 from 2 to 1
		temp := strings.Split(row, " ")
		count, err := strconv.Atoi(temp[1])
		if err != nil {
			panic(err)
		}
		from, err := strconv.Atoi(temp[3])
		if err != nil {
			panic(err)
		}
		to, err := strconv.Atoi(temp[5])
		if err != nil {
			panic(err)
		}
		for i := 0; i < count; i++ {
			temp, err := stacks[from-1].Pop()
			if err != nil {
				panic(err)
			}
			stacks[to-1].Push(temp)
		}
	}

	ret := ""

	for _, s := range stacks {
		val, err := s.Pop()
		if err != nil {
			panic(err)
		}
		ret += val
	}

	return ret
}

func Day05part02(input string) string {
	temp := strings.Split(input, "\n\n")

	rows := strings.Split(temp[0], "\n")

	stacks := make([]Stack, (len(rows[0])/4)+1)

	for i := len(rows) - 1; i >= 0; i-- {
		content := SplitN(rows[i], 4)
		for j, c := range content {
			if c[2] != ' ' {
				stacks[j].Push(string(c[1]))
			}
		}
	}

	rows = strings.Split(temp[1], "\n")

	for _, row := range rows {
		// move 1 from 2 to 1
		temp := strings.Split(row, " ")
		count, err := strconv.Atoi(temp[1])
		if err != nil {
			panic(err)
		}
		from, err := strconv.Atoi(temp[3])
		if err != nil {
			panic(err)
		}
		to, err := strconv.Atoi(temp[5])
		if err != nil {
			panic(err)
		}
		var tempStack Stack

		for i := 0; i < count; i++ {
			temp, err := stacks[from-1].Pop()
			if err != nil {
				panic(err)
			}
			tempStack.Push(temp)
		}

		for i := 0; i < count; i++ {
			temp, err := tempStack.Pop()
			if err != nil {
				panic(err)
			}
			stacks[to-1].Push(temp)
		}
	}

	ret := ""

	for _, s := range stacks {
		val, err := s.Pop()
		if err != nil {
			panic(err)
		}
		ret += val
	}

	return ret
}
