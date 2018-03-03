package undergo_test

import (
	"testing"
	. "undergo"
)

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello Go!"

	if got != want {
		t.Error(got, "!=", want)
	}
}
