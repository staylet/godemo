/*
A Test Go Public module.
*/

package demo

import (
	"testing"
)

func TestHello(t *testing.T) {
	cases := []struct {
		want string
	}{
		{"Hello, Gopher!"},
	}
	for _, c := range cases {
		got := Hello()
		if got != c.want {
			t.Errorf("Hello() = %#v, want %#v", got, c.want)
		}
	}
}

func TestGoodbye(t *testing.T) {
	cases := []struct {
		want string
	}{
		{"Goodbye, Gopher!"},
	}
	for _, c := range cases {
		got := Goodbye()
		if got != c.want {
			t.Errorf("Goodbye() = %#v, want %#v", got, c.want)
		}
	}
}
