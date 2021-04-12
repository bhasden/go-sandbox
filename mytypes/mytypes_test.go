package mytypes_test

import (
	"sandbox/mytypes"
	"testing"
)

func TestMyStringLen(t *testing.T) {
	s := mytypes.MyString("Hello, Gophers!")
	want := 15
	got := s.Len()
	if want != got {
		t.Errorf("Wrong len of string, expected: %d, got %d", want, got)
	}
}

func TestMultiStringJoin(t *testing.T) {
	m := mytypes.MyMultiString{"a", "b", "c"}
	want := "a plus b plus c"
	got := m.Join(" plus ")
	if want != got {
		t.Errorf("Wrong value for Join, expected: %s, got %s", want, got)
	}
}
