package main

import "testing"

func TestHelloWorld(t *testing.T) {
	cases := []struct {
		want string
	}{
		{"Hello World!"},
	}

	for _, c := range cases {
		got := HelloWorld()
		if got != c.want {
			t.Errorf("HelloWorld() == %q, want %q", got, c.want)
		}
	}
}
