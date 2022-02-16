package iterrors

import (
	"errors"
	"testing"
)

func TestWrap(t *testing.T) {
	var err error
	Wrap(&err, "whatever")
	if err != nil {
		t.Errorf("got %v, want nil", err)
	}

	orig := errors.New("bad stuff")
	err = orig
	want := "Frob(3): bad stuff"
	if got := err.Error(); got != want {
		t.Errorf("got %v, want %s", got, want)
	}
	if got := errors.Unwrap(err); got != orig {
		t.Errorf("Unwrap: got %#v, want %#v", got, orig)
	}
}
