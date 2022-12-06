package main

import (
	"os"
	"testing"
)

func getTestDataDay06part01() struct {
	in   string
	want string
} {
	bdata, err := os.ReadFile("test/Day06.data")
	if err != nil {
		panic(err)
	}
	bans, err := os.ReadFile("test/Day06part01.answer")
	if err != nil {
		panic(err)
	}
	temp := struct {
		in   string
		want string
	}{
		string(bdata),
		string(bans),
	}

	return temp
}

func getTestDataDay06part02() struct {
	in   string
	want string
} {
	bdata, err := os.ReadFile("test/Day06.data")
	if err != nil {
		panic(err)
	}
	bans, err := os.ReadFile("test/Day06part02.answer")
	if err != nil {
		panic(err)
	}
	temp := struct {
		in   string
		want string
	}{
		string(bdata),
		string(bans),
	}

	return temp
}

func TestDay06part01(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
		{"nppdvjthqldpwncqszvftbrmjlhg", 6},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
	}

	for _, c := range cases {
		got := Day06part01(c.in)
		if got != c.want {
			t.Errorf("Day06(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}

func TestDay06part02(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 23},
		{"nppdvjthqldpwncqszvftbrmjlhg", 23},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26},
	}

	for _, c := range cases {
		got := Day06part02(c.in)
		if got != c.want {
			t.Errorf("Day06(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}
