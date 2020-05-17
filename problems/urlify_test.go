package problems

import "testing"

func TestUrlify(t *testing.T) {
	s := []byte("ab c  ")
	len := 4
	actual := Urlify(s, len)
	expected := []byte("ab%20c")

	expectedS := string(expected)
	actualS := string(actual)

	if actualS != expectedS {
		t.Errorf("Expect %s, actual %s", string(expected), string(actual))
	}
}
