package undergo_test

import (
	"testing"
	. "github.com/stsh89/undergo"
)

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello Go!"

	if got != want {
		t.Error(got, "!=", want)
	}
}
