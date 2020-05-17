package problems

import "testing"

func TestIsUnique(t *testing.T) {
	actual := IsUnique("abca")
	expected := false

	if actual != expected {
		t.Errorf("Expect %t, actual %t", expected, actual)
	}
}

func TestIsUniqueUtf8(t *testing.T) {
	actual := IsUnique("фБгїф")
	expected := false

	if actual != expected {
		t.Errorf("Expect %t, actual %t", expected, actual)
	}
}
