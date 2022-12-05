package main

import (
	"os"
	"strconv"
	"testing"
)

func getTestDataDay03part01() struct {
	in   string
	want int
} {
	bdata, err := os.ReadFile("test/Day03.data")
	if err != nil {
		panic(err)
	}
	bans, err := os.ReadFile("test/Day03part01.answer")
	if err != nil {
		panic(err)
	}
	ans, err := strconv.Atoi(string(bans))
	if err != nil {
		panic(err)
	}
	temp := struct {
		in   string
		want int
	}{
		string(bdata),
		ans,
	}

	return temp
}

func getTestDataDay03part02() struct {
	in   string
	want int
} {
	bdata, err := os.ReadFile("test/Day03.data")
	if err != nil {
		panic(err)
	}
	bans, err := os.ReadFile("test/Day03part02.answer")
	if err != nil {
		panic(err)
	}
	ans, err := strconv.Atoi(string(bans))
	if err != nil {
		panic(err)
	}
	temp := struct {
		in   string
		want int
	}{
		string(bdata),
		ans,
	}

	return temp
}

func TestDay03part01(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"vJrwpWtwJgWrhcsFMMfFFhFp", 16},
		{"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", 38},
		{"PmmdzqPrVvPwwTWBwg", 42},
		{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", 22},
		{"ttgJtRGJQctTZtZT", 20},
		{"CrZsJsPPZsGzwwsLwLmpwMDw", 19},
		getTestDataDay03part01(),
	}

	for _, c := range cases {
		got := Day03part01(c.in)
		if got != c.want {
			t.Errorf("Day03(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}

func TestDay03part02(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"vJrwpWtwJgWrhcsFMMfFFhFp\njqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\nPmmdzqPrVvPwwTWBwg", 18},
		{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn\nttgJtRGJQctTZtZT\nCrZsJsPPZsGzwwsLwLmpwMDw", 52},
		getTestDataDay03part02(),
	}

	for _, c := range cases {
		got := Day03part02(c.in)
		if got != c.want {
			t.Errorf("Day03(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}
