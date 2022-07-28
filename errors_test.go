package errors

import (
	"fmt"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	want := "foo"
	got := New(want)

	if got == nil || got.Error() != want {
		t.Errorf("New not done, got: %s want: %s", got.Error(), want)
	}
}

func TestErrorf(t *testing.T) {
	want := "foo bar 1"
	got := Errorf("foo %s %d", "bar", 1)

	if got == nil || got.Error() != want {
		t.Errorf("Errorf not done, got: %s want: %s", got.Error(), want)
	}
}

func TestWrapNil(t *testing.T) {
	got := Wrap(nil, "null")

	if got != nil {
		t.Errorf("Wrap with err is nil not done, got: %#v want: nil", got)
	}
}

func TestWrap(t *testing.T) {
	want := "bar: foo"
	got := Wrap(New("foo"), "bar")

	if got == nil || got.Error() != want {
		t.Errorf("Wrap not done, got: %s want: %s", got.Error(), want)
	}
}

func TestWrapfNil(t *testing.T) {
	got := Wrapf(nil, "null")

	if got != nil {
		t.Errorf("Wrapf with err is nil not done, got: %#v want: nil", got)
	}
}

func TestWrapf(t *testing.T) {
	want := "bar 1: foo"
	got := Wrapf(New("foo"), "bar %d", 1)

	if got == nil || got.Error() != want {
		t.Errorf("Wrapf not done, got: %s want: %s", got.Error(), want)
	}
}

func TestCause(t *testing.T) {
	want := "foo"

	got := fmt.Errorf(want)
	got = Wrap(got, "bar")
	got = Wrapf(got, "baz %d", 1)

	if got == nil || Cause(got).Error() != want {
		t.Errorf("Cause not done, got: %s want: %s", Cause(got).Error(), want)
	}
}

func TestSetSeparator(t *testing.T) {
	sep := "-> "

	SetSeparator(sep)

	err := New("foo")
	err = Wrap(err, "bar")

	if !strings.Contains(err.Error(), sep) {
		t.Errorf("SetSeparator not done, err=%s", err.Error())
	}
}
