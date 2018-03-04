package hello_test

import (
	. "github.com/stsh89/undergo/hello"
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello Go!"

	if got != want {
		t.Error(got, "!=", want)
	}
}
