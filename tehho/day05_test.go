package main

import (
	"os"
	"testing"
)

func getTestDataDay05part01() struct {
	in   string
	want string
} {
	bdata, err := os.ReadFile("test/Day05.data")
	if err != nil {
		panic(err)
	}
	bans, err := os.ReadFile("test/Day05part01.answer")
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

func getTestDataDay05part02() struct {
	in   string
	want string
} {
	bdata, err := os.ReadFile("test/Day05.data")
	if err != nil {
		panic(err)
	}
	bans, err := os.ReadFile("test/Day05part02.answer")
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

func TestDay05part01(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		getTestDataDay05part01(),
	}

	for _, c := range cases {
		got := Day05part01(c.in)
		if got != c.want {
			t.Errorf("Day05(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestDay05part02(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		getTestDataDay05part02(),
	}

	for _, c := range cases {
		got := Day05part02(c.in)
		if got != c.want {
			t.Errorf("Day05(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
