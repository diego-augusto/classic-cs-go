package fib

import (
	"testing"
)

func TestFib(t *testing.T) {

	/*
	 * p: 0 1 2 3 4 5 6 7
	 * s: 0 1 1 2 3 5 8 13...
	 */

	cases := []struct {
		p int
		e int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
		{50, 12586269025},
	}

	for _, c := range cases {
		r := fib(c.p)
		if r != c.e {
			t.Errorf("\ninput: %d\ngot: %d\nwant: %d", c.p, r, c.e)
		}
	}
}
