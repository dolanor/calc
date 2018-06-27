package calc_test

import (
	"testing"

	"github.com/dolanor/calc"
)

func TestAddOK(t *testing.T) {
	cases := []struct {
		a, b     int
		expected int
	}{
		{1, 2, 3},
		{2, 2, 4},
	}

	for _, c := range cases {
		got := calc.Add(c.a, c.b)
		if got != c.expected {
			t.Errorf("we expected %d but got %d", c.expected, got)
		}
	}
}

func TestDivOK(t *testing.T) {
	cases := []struct {
		a, b     int
		expected int
	}{
		{1, 2, 0},
		{2, 2, 1},
	}

	for _, c := range cases {
		got, err := calc.Div(c.a, c.b)
		if err != nil {
			t.Errorf("got error %v", err)
		}
		if got != c.expected {
			t.Errorf("we expected %d but got %d", c.expected, got)
		}
	}
}

func TestDivError(t *testing.T) {
	cases := []struct {
		a, b   int
		gotErr bool
	}{
		{1, 2, false},
		{2, 2, false},
		{2, 0, true},
	}

	for _, c := range cases {
		_, err := calc.Div(c.a, c.b)
		if (err != nil) != c.gotErr {
			t.Errorf("got error %v", err)
		}
	}
}
