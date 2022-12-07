package main

import (
	"os"
	"strconv"
	"testing"
)

func getTestDataDay07part01() struct {
	in   string
	want int
} {
	bdata, err := os.ReadFile("test/Day07.data")
	if err != nil {
		panic(err)
	}
	bans, err := os.ReadFile("test/Day07part01.answer")
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

func getTestDataDay07part02() struct {
	in   string
	want int
} {
	bdata, err := os.ReadFile("test/Day07.data")
	if err != nil {
		panic(err)
	}
	bans, err := os.ReadFile("test/Day07part02.answer")
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

func TestDay07part01(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		getTestDataDay07part01(),
	}

	for _, c := range cases {
		got := Day07part01(c.in)
		if got != c.want {
			t.Errorf("Day07(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}

func TestDay07part02(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		getTestDataDay07part02(),
	}

	for _, c := range cases {
		got := Day07part02(c.in)
		if got != c.want {
			t.Errorf("Day07(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}
