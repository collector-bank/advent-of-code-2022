package main

import (
	"os"
	"strconv"
	"testing"
)

func getTestDataDay04part01() struct {
	in   string
	want int
} {
	bdata, err := os.ReadFile("test/Day04.data")
	if err != nil {
		panic(err)
	}
	bans, err := os.ReadFile("test/Day04part01.answer")
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

func getTestDataDay04part02() struct {
	in   string
	want int
} {
	bdata, err := os.ReadFile("test/Day04.data")
	if err != nil {
		panic(err)
	}
	bans, err := os.ReadFile("test/Day04part02.answer")
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

func TestDay04part01(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		getTestDataDay04part01(),
	}

	for _, c := range cases {
		got := Day04part01(c.in)
		if got != c.want {
			t.Errorf("Day04(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}

func TestDay04part02(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"6-6,5-7", 1},
		{"5-7,6-6", 1},
		{"5-7,6-8", 1},
		{"5-6,7-8", 0},
		getTestDataDay04part02(),
	}

	for _, c := range cases {
		got := Day04part02(c.in)
		if got != c.want {
			t.Errorf("Day04(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}
