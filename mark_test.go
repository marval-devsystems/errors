package errors

import "testing"

func TestMarkNil(t *testing.T) {
	got := Mark(nil, 0)

	if got != nil {
		t.Errorf("Mark with err is nil not done, got: %#v want: nil", got)
	}
}

func TestMark(t *testing.T) {
	want := uint32(1000)

	err := New("foo")

	err = Mark(err, want)

	code, ok := GetMark(err)
	if !ok || code != want {
		t.Errorf("Mark not done, got: %d want: %d", code, want)
	}
}

func TestMarkfNil(t *testing.T) {
	got := Markf(nil, 0, "")

	if got != nil {
		t.Errorf("Markf with err is nil not done, got: %#v want: nil", got)
	}
}

func TestMarkf(t *testing.T) {
	separator = ": "

	want := uint32(1000)
	wantStr := "foo bar: foo"

	err := New("foo")

	err = Markf(err, want, "foo %s", "bar")

	code, ok := GetMark(err)
	if !ok || code != want || err.Error() != wantStr {
		t.Errorf("Markf not done, got code: %d want code: %d, got err: %s want err: %s", code, want, err.Error(), wantStr)
	}
}
