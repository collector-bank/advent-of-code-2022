package main

import (
	"os"
	"strconv"
	"testing"
)

func getTestDataDay01part01() struct {
	in   string
	want int
} {
	bdata, err := os.ReadFile("test/Day01.data")
	if err != nil {
		panic(err)
	}
	bans, err := os.ReadFile("test/Day01part01.answer")
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

func getTestDataDay01part02() struct {
	in   string
	want int
} {
	bdata, err := os.ReadFile("test/Day01.data")
	if err != nil {
		panic(err)
	}
	bans, err := os.ReadFile("test/Day01part02.answer")
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

func TestDay01part01(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"1000", 1000},
		getTestDataDay01part01(),
	}

	for _, c := range cases {
		got := Day01part01(c.in)
		if got != c.want {
			t.Errorf("Day01(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}

func TestDay01part02(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"1000", 1000},
		getTestDataDay01part02(),
	}

	for _, c := range cases {
		got := Day01part02(c.in)
		if got != c.want {
			t.Errorf("Day01(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}
