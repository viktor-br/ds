package problems

import "testing"

func TestSingleTransformation(t *testing.T) {
	result := SingleTransformation("pepa", "ppa")

	if !result {
		t.Error("Expect single transformation true")
	}

	result = SingleTransformation("pepa", "pepa")

	if !result {
		t.Error("Expect single transformation true")
	}

	result = SingleTransformation("pea", "pepa")

	if !result {
		t.Error("Expect single transformation true")
	}

	result = SingleTransformation("pepa", "test")

	if result {
		t.Error("Expect single transformation true")
	}
}
