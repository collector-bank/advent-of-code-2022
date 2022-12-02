package main

import (
	"os"
	"strconv"
	"testing"
)

func getTestDataDay02part01() struct {
	in   string
	want int
} {
	bdata, err := os.ReadFile("test/Day02.data")
	if err != nil {
		panic(err)
	}
	bans, err := os.ReadFile("test/Day02part01.answer")
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

func getTestDataDay02part02() struct {
	in   string
	want int
} {
	bdata, err := os.ReadFile("test/Day02.data")
	if err != nil {
		panic(err)
	}
	bans, err := os.ReadFile("test/Day02part02.answer")
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

func TestDay02part01(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"A Y", 8},
		{"B X", 1},
		{"C Z", 6},
		{"A Y\nB X\nC Z", 15},
		getTestDataDay02part01(),
	}

	for _, c := range cases {
		got := Day02part01(c.in)
		if got != c.want {
			t.Errorf("Day02(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}

func TestDay02part02(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"A Y", 4},
		{"B X", 1},
		{"C Z", 7},
		getTestDataDay02part02(),
	}

	for _, c := range cases {
		got := Day02part02(c.in)
		if got != c.want {
			t.Errorf("Day02(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}
